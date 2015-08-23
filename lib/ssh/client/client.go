package client

import (
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
	"net"
	"os"
)

// Initialize SSH client (our transport backbone)
func New(
	sshServer string,
	sshUsername string,
	sshPassword string,
	sshUseAgent bool,
) (*ssh.Client, error) {
	sshAuth, err := getSSHAuth(sshUseAgent, sshPassword)
	if err != nil {
		return nil, err
	}

	sshClientConfig := &ssh.ClientConfig{
		User: sshUsername,
		Auth: sshAuth,
	}

	sshClient, err := ssh.Dial("tcp", sshServer, sshClientConfig)
	if err != nil {
		return nil, err
	}

	return sshClient, nil
}

func getSSHAuth(sshUseAgent bool, sshPassword string) ([]ssh.AuthMethod, error) {
	if !sshUseAgent {
		return []ssh.AuthMethod{ssh.Password(sshPassword)}, nil
	}

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
