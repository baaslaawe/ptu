package arguments

import (
	"errors"
	"flag"
	"os"
)

// No arguments passed? :(
func IsListEmpty() bool {
	return len(os.Args) < 2
}

// Is help requested by passing -h|--help as an argument?
func IsHelpRequested() bool {
	return helpArgumentRegexp.MatchString(os.Args[1])
}

// Parse command line arguments, perform some initial validation and variable mutation
func Parse() (*Config, error) {
	var sshServer = flag.String("s", defaultSSHServer, "SSH server (host[:port]) to connect")
	var sshUsername = flag.String("u", defaultSSHUsername, "username to connect SSH server")
	var sshPassword = flag.String("p", defaultSSHPassword, "password to authenticate against SSH server (do not use, please)")
	var targetHost = flag.String("t", defaultTargetHost, "target host:port we will forward connections to")
	var exposedBind = flag.String("b", defaultExposedBind, "bind (listener) to expose on the SSH server side")
	var exposedPort = flag.Int("e", defaultExposedPort, "port to expose and forward on the SSH server side")

	flag.Parse()

	if !isSSHServerSet(*sshServer) {
		return nil, errors.New("SSH server not defined")
	}

	if !isTCPPortValid(*exposedPort) {
		return nil, errors.New("exposed TCP port number is invalid")
	}

	if !isHostWithPort(*sshServer) {
		*sshServer = concatHostPort(*sshServer, defaultSSHPort)
	}

	if !isHostWithPort(*targetHost) {
		*targetHost = concatHostPort(*targetHost, defaultTargetPort)
	}

	config := &Config{
		SSHServer:   *sshServer,
		SSHUsername: *sshUsername,
		SSHPassword: *sshPassword,
		SSHUseAgent: !isSSHPasswordSet(*sshPassword),
		TargetHost:  *targetHost,
		ExposedBind: *exposedBind,
		ExposedPort: *exposedPort,
		ExposedHost: concatHostPort(*exposedBind, *exposedPort),
		ConnectTo:   mergeHostPort(*sshServer, *exposedPort),
	}

	return config, nil
}
