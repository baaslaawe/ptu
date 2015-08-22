package client

import (
	"../../util/config"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
	"net"
	"os"
)

// Initialize SSH client (our transport backbone)
func New(config config.Config) (*ssh.Client, error) {
	sshAuth, err := getSSHAuth(config)
	if err != nil {
		return nil, err
	}

	sshClientConfig := &ssh.ClientConfig{
		User: config.SSHUsername,
		Auth: sshAuth,
	}

	sshClient, err := ssh.Dial("tcp", config.SSHServer, sshClientConfig)
	if err != nil {
		return nil, err
	}

	return sshClient, nil
}

func getSSHAuth(config config.Config) ([]ssh.AuthMethod, error) {
	if config.SSHUseAgent {
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
	}

	return []ssh.AuthMethod{ssh.Password(config.SSHPassword)}, nil
}
