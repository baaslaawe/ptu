package config

import (
	"errors"
	"flag"
	"net"
	"os"
	"regexp"
	"strconv"
)

// Config{} is a container for ptu configuration
type Config struct {
	SSHServer   string `yaml:"s"`
	SSHUsername string `yaml:"u"`
	SSHPassword string `yaml:"p"`
	SSHUseAgent bool
	TargetHost  string `yaml:"t"`
	ExposedBind string `yaml:"b"`
	ExposedPort int    `yaml:"e"`
	ExposedHost string
	ConnectTo   string
}

// IsHelpRequested() checks, if help was requested (by passing -h|--help as an argument)
func IsHelpRequested() bool {
	var helpArgumentRegexp = regexp.MustCompile(`^(-h|--help)$`)

	if len(os.Args) < 2 {
		return false
	}

	return helpArgumentRegexp.MatchString(os.Args[1])
}

// ParseArguments() parses command line arguments, performs some initial validation and variable mutation
func ParseArguments(d *Config) (*Config, error) {
	var sshServer = flag.String("s", d.SSHServer, "SSH server (host[:port]) to connect")
	var sshUsername = flag.String("u", d.SSHUsername, "username to connect SSH server")
	var sshPassword = flag.String("p", d.SSHPassword, "password to authenticate against SSH server")
	var targetHost = flag.String("t", d.TargetHost, "target host:port we will forward connections to")
	var exposedBind = flag.String("b", d.ExposedBind, "bind (listener) to expose on the SSH server side")
	var exposedPort = flag.Int("e", d.ExposedPort, "port to expose and forward on the SSH server side")

	flag.Parse()

	if !isSSHServerSet(*sshServer) {
		return nil, errors.New("SSH server not defined (try to run program with `--help` option)")
	}

	if !isTCPPortValid(*exposedPort) {
		return nil, errors.New("exposed TCP port number is invalid")
	}

	if !isHostWithPort(*sshServer) {
		*sshServer = joinHostPort(*sshServer, defaultSSHPort)
	}

	if !isHostWithPort(*targetHost) {
		*targetHost = joinHostPort(*targetHost, defaultTargetPort)
	}

	config := &Config{
		SSHServer:   *sshServer,
		SSHUsername: *sshUsername,
		SSHPassword: *sshPassword,
		SSHUseAgent: !isSSHPasswordSet(*sshPassword),
		TargetHost:  *targetHost,
		ExposedBind: *exposedBind,
		ExposedPort: *exposedPort,
		ExposedHost: joinHostPort(*exposedBind, *exposedPort),
		ConnectTo:   mergeHostPort(*sshServer, *exposedPort),
	}

	return config, nil
}

func isSSHServerSet(s string) bool { return s != "" }

func isSSHPasswordSet(s string) bool { return s != "" }

func isTCPPortValid(port int) bool { return !(port < 1) || (port > 65535) }

func isHostWithPort(hostPort string) bool {
	_, _, err := net.SplitHostPort(hostPort)

	return err == nil
}

func joinHostPort(host string, port int) string {
	return net.JoinHostPort(host, strconv.Itoa(port))
}

func mergeHostPort(hostPort string, port int) string {
	host, _, _ := net.SplitHostPort(hostPort)

	return joinHostPort(host, port)
}
