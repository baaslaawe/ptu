package main

import (
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
	"net"
	"os"
	"strconv"
	"strings"
)

// Initialize SSH listener 0.0.0.0:<forward_port> for connection forwarding
func InitializeSSHListener(sshUsername string, sshServer string, exposePort int) (net.Listener, error) {
	sshAgentConn, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK"))
	if err != nil {
		return nil, err
	}

	sshAgent := agent.NewClient(sshAgentConn)
	sshAgentSigners, err := sshAgent.Signers()
	if err != nil {
		return nil, err
	}

	sshAuth := []ssh.AuthMethod{ssh.PublicKeys(sshAgentSigners...)}
	sshClientConfig := &ssh.ClientConfig{
		User: sshUsername,
		Auth: sshAuth,
	}
	sshClient, err := ssh.Dial("tcp", sshServer, sshClientConfig)
	if err != nil {
		return nil, err
	}

	forwardListenTo := strings.Join([]string{"0.0.0.0", strconv.Itoa(exposePort)}, ":")
	sshListener, err := sshClient.Listen("tcp", forwardListenTo)
	if err != nil {
		return nil, err
	}

	return sshListener, nil
}
