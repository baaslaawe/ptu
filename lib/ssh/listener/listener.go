package listener

import (
	"golang.org/x/crypto/ssh"
	"net"
)

// New sets up a new instance of the <exposed_bind>:<exposed_port> listener
func New(sshClient *ssh.Client, exposedHost string) (net.Listener, error) {
	sshListener, err := sshClient.Listen("tcp", exposedHost)
	if err != nil {
		return nil, err
	}

	return sshListener, nil
}
