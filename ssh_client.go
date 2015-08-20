package main

import (
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
	"net"
	"os"
)

// Initialize SSH listener 0.0.0.0:<forward_port> for connection forwarding
func InitSSHClient(config Config) (*ssh.Client, error) {
	sshAuth, err := getSSHAuth(config)
	if err != nil {
		return nil, err
	}

	sshClientConfig := &ssh.ClientConfig{
		User: config.sshUsername,
		Auth: sshAuth,
	}
	sshClient, err := ssh.Dial("tcp", config.sshServer, sshClientConfig)
	if err != nil {
		return nil, err
	}

	return sshClient, nil
}

func getSSHAuth(config Config) ([]ssh.AuthMethod, error) {
	if config.sshUseAgent {
		sshAgentConn, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK"))
		if err != nil {
			return nil, err
		}

		sshAgent := agent.NewClient(sshAgentConn)

		sshAgentSigners, err := sshAgent.Signers()
		if err != nil {
			return nil, err
		}

		return []ssh.AuthMethod{ssh.PublicKeys(sshAgentSigners...)}, nil
	} else {
		return []ssh.AuthMethod{ssh.Password(config.sshPassword)}, nil
	}
}
