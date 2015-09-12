package forwarder

import (
	"io"
	"log"
	"net"
)

// Forward forwards connections from SSH listener to the target host
func Forward(sshListener net.Listener, targetHost string) error {
	sshConn, err := sshListener.Accept()
	if err != nil {
		log.Printf("Error accepting connection on SSH listener: %s", err)

		return err
	}

	targetConn, err := net.Dial("tcp", targetHost)
	if err != nil {
		log.Printf("Error establishing remote host connection: %s", err)

		return nil
	}

	log.Printf(
		"[CONN] From: %v | Forward: %v => %v",
		sshConn.RemoteAddr(),
		targetConn.LocalAddr(),
		targetConn.RemoteAddr(),
	)

	go func() {
		_, err = io.Copy(sshConn, targetConn)
		if err != nil {
			log.Printf(
				"ssh => target IO error: %s (%v => %v)",
				err,
				sshConn.RemoteAddr(),
				targetConn.RemoteAddr(),
			)

			return
		}
	}()

	go func() {
		_, err = io.Copy(targetConn, sshConn)
		if err != nil {
			log.Printf(
				"target => ssh IO error: %s (%v => %v)",
				err,
				targetConn.RemoteAddr(),
				sshConn.RemoteAddr(),
			)

			return
		}
	}()

	return nil
}
