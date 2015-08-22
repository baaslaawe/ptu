package listener

import (
	"golang.org/x/crypto/ssh"
	"net"
)

// Set up SSH listener <exposed_host>:<exposed_port> for connection forwarding
func New(sshClient *ssh.Client, exposedBind string) (net.Listener, error) {
	sshListener, err := sshClient.Listen("tcp", exposedBind)
	if err != nil {
		return nil, err
	}

	return sshListener, nil
}
