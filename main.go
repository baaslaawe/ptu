package main

import (
	"./lib/net/forwarder"
	"./lib/ssh/client"
	"./lib/ssh/listener"
	"./lib/util/arguments"
	"./lib/util/display"
	"log"
)

func main() {
	// Some tender erotic foreplay
	if arguments.IsListEmpty() || arguments.IsHelpRequested() {
		display.PrintHelpAndExit()
	}

	c, err := arguments.Parse()
	if err != nil {
		log.Fatalf("Error while parsing command line arguments: %s", err)
	}

	display.PrintGatewayPortsNB()

	display.PrintConfig(c.SSHServer, c.SSHUsername, c.SSHUseAgent, c.TargetHost, c.ConnectTo)

	// Initialize SSH client
	sshClient, err := client.New(c.SSHServer, c.SSHUsername, c.SSHPassword, c.SSHUseAgent)
	if err != nil {
		log.Fatalf("Error initializing SSH client %s", err)
	}

	// Set up SSH listener <exposed_bind>:<exposed_port> on the SSH server <ssh_server>:<ssh_port>
	sshListener, err := listener.New(sshClient, c.ExposedHost)
	if err != nil {
		log.Fatalf("Error setting up SSH listener %s", err)
	}

	// Vamos muchachos!
	for {
		forwarder.Forward(sshListener, c.TargetHost)
	}
}
