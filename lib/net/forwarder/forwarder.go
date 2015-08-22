package forwarder

import (
	"io"
	"log"
	"net"
)

// Forward connection from SSH listener to remote host
func Forward(sshListener net.Listener, targetHost string) {
	sshConn, err := sshListener.Accept()
	if err != nil {
		log.Printf("Error accepting connection on SSH listener: %s", err)
	}

	remoteConn, err := net.Dial("tcp", targetHost)
	if err != nil {
		log.Printf("Error establishing remote host connection: %s", err)
	}

	log.Printf(
		"[CONN] SSH: %v => %v | HOST: %v => %v",
		sshConn.RemoteAddr(),
		sshConn.LocalAddr(),
		remoteConn.LocalAddr(),
		remoteConn.RemoteAddr(),
	)

	go func() {
		_, err = io.Copy(sshConn, remoteConn)
		if err != nil {
			log.Printf(
				"> ssh => remote IO error: %s (%v => %v)",
				err,
				sshConn.RemoteAddr(),
				remoteConn.RemoteAddr(),
			)
		}
	}()

	go func() {
		_, err = io.Copy(remoteConn, sshConn)
		if err != nil {
			log.Printf(
				"< remote => ssh IO error: %s (%v => %v)",
				err,
				remoteConn.RemoteAddr(),
				sshConn.RemoteAddr(),
			)
		}
	}()
}
