package config

import (
	"errors"
	"flag"
	"net"
	"os"
	"regexp"
	"strconv"
)

// Config is a container for ptu configuration
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

// IsHelpRequested checks, if help was requested (by passing -h|--help as an argument)
func IsHelpRequested() bool {
	var helpArgumentRegexp = regexp.MustCompile(`^(-h|--help)$`)

	if len(os.Args) < 2 {
		return false
	}

	return helpArgumentRegexp.MatchString(os.Args[1])
}

// ParseArguments parses command line arguments, performs some initial validation and variable mutation
func ParseArguments(d *Config) (*Config, error) {
	var fs = flag.String("s", d.SSHServer, "SSH server (host[:port]) to connect")
	var fu = flag.String("u", d.SSHUsername, "username to connect SSH server")
	var fp = flag.String("p", d.SSHPassword, "password to authenticate against SSH server")
	var ft = flag.String("t", d.TargetHost, "target host:port we will forward connections to")
	var fb = flag.String("b", d.ExposedBind, "bind (listener) to expose on the SSH server side")
	var fe = flag.Int("e", d.ExposedPort, "port to expose and forward on the SSH server side")

	var YAMLConfig = flag.String("c", "", "YAML config name to load from '~/.ptu' directory")

	flag.Parse()

	c := &Config{
		SSHServer:   *fs,
		SSHUsername: *fu,
		SSHPassword: *fp,
		TargetHost:  *ft,
		ExposedBind: *fb,
		ExposedPort: *fe,
	}

	if isStringParamSet(*YAMLConfig) {
		c, err := applyYAMLConfig(*YAMLConfig, c)
		if err != nil {
			return nil, err
		}

		// This is required to override YAML config settings with ones passed by arguments
		flag.Visit(func(f *flag.Flag) {
			switch f.Name {
			case "s":
				c.SSHServer = f.Value.String()
			case "u":
				c.SSHUsername = f.Value.String()

			case "p":
				c.SSHPassword = f.Value.String()

			case "t":
				c.TargetHost = f.Value.String()

			case "b":
				c.ExposedBind = f.Value.String()

			case "e":
				e, _ := strconv.Atoi(f.Value.String())
				c.ExposedPort = e
			}
		})
	}

	c, err := prepareConfig(c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func applyYAMLConfig(YAMLConfig string, c *Config) (*Config, error) {
	if !isStringParamSet(YAMLConfig) {
		return c, nil
	}

	a, err := LoadYAML(YAMLConfig, c)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func prepareConfig(c *Config) (*Config, error) {
	if !isStringParamSet(c.SSHServer) {
		return nil, errors.New("SSH server not defined (try to run program with `--help` option)")
	}

	if !isTCPPortValid(c.ExposedPort) {
		return nil, errors.New("exposed TCP port number is invalid")
	}

	if !isHostWithPort(c.SSHServer) {
		c.SSHServer = joinHostPort(c.SSHServer, defaultSSHPort)
	}

	if !isHostWithPort(c.TargetHost) {
		c.TargetHost = joinHostPort(c.TargetHost, defaultTargetPort)
	}

	c.SSHUseAgent = !isStringParamSet(c.SSHPassword)
	c.ExposedHost = joinHostPort(c.ExposedBind, c.ExposedPort)

	c.ConnectTo = mergeHostPort(c.SSHServer, c.ExposedPort)

	return c, nil
}

func isStringParamSet(s string) bool { return s != "" }

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
