package main

import (
	"./lib/net/forwarder"
	"./lib/ssh/client"
	"./lib/ssh/listener"
	"./lib/util/config"
	"./lib/util/display"

	"log"
	"time"
)

const retrySeconds = 10

func main() {
	// Some tender erotic foreplay
	if config.IsHelpRequested() {
		display.PrintHelpAndExit()
	}

	// Load defaults: built-in or from file, if it exists
	d, err := config.LoadDefaults()
	if err != nil {
		log.Fatalf("Unable to load defaults: %s", err)
	}

	// Merge default config with params passed as command line arguments
	c, err := config.ParseArguments(d)
	if err != nil {
		log.Fatalf("Error while parsing command line arguments: %s", err)
	}

	display.PrintGatewayPortsNB()

	display.PrintConfig(c.SSHServer, c.SSHUsername, c.SSHUseAgent, c.TargetHost, c.ConnectTo)

RETRY:
	// Initialize SSH client (at least try to!)
	sshClient, err := client.New(c.SSHServer, c.SSHUsername, c.SSHPassword, c.SSHUseAgent)
	if err != nil {
		log.Printf("Error initializing SSH client: %s (will retry)", err)
		time.Sleep(retrySeconds * time.Second)
		goto RETRY
	} else {
		log.Printf("[OK] SSH client initialized!")
	}

	// Set up SSH listener <exposed_bind>:<exposed_port> on the SSH server <ssh_server>:<ssh_port>
	sshListener, err := listener.New(sshClient, c.ExposedHost)
	if err != nil {
		log.Fatalf("Error setting up SSH listener: %s", err)
	}

	// Vamos muchachos!
	for {
		err := forwarder.Forward(sshListener, c.TargetHost)
		if err != nil {
			log.Printf("[OMG] Critical failure on forwarder! Will re-setup SSH connection.")
			time.Sleep(retrySeconds * time.Second)
			goto RETRY
		}
	}
}
