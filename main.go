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

	config, err := arguments.Parse()
	if err != nil {
		log.Fatalf("Error while parsing command line arguments: %s", err)
	}

	display.PrintGatewayPortsNB()

	display.PrintConfig(*config)

	// Initialize SSH client
	sshClient, err := client.New(*config)
	if err != nil {
		log.Fatalf("Error initializing SSH client %s", err)
	}

	// Set up SSH listener <exposed_bind>:<exposed_port> on the SSH server <ssh_server>:<ssh_port>
	sshListener, err := listener.New(sshClient, config.ExposedHost)
	if err != nil {
		log.Fatalf("Error setting up SSH listener %s", err)
	}

	// Vamos muchachos!
	for {
		forwarder.Forward(sshListener, config.TargetHost)
	}
}
