package main

import (
	"log"
	"os"
)

func main() {
	// Some tender erotic foreplay
	if len(os.Args) < 2 || IsHelpRequested(os.Args[1]) {
		DisplayHelp()
		os.Exit(1)
	}

	config, err := ParseCommandLineArguments()
	if err != nil {
		log.Fatalf("Error while parsing command line arguments: %s", err)
	}

	DisplayNB()

	DisplayConfig(*config)

	// Initialize SSH client
	sshClient, err := InitSSHClient(*config)
	if err != nil {
		log.Fatalf("Error initializing SSH client %s", err)
	}

	// Set up SSH listener <exposed_host>:<exposed_port> on the SSH server <ssh_server>:<ssh_port>
	sshListener, err := SetupSSHListener(sshClient, config.exposedBind)
	if err != nil {
		log.Fatalf("Error setting up SSH listener %s", err)
	}

	// Vamos muchachos!
	for {
		Forward(sshListener, config.targetHost)
	}
}
