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

	sshServer, targetHost, exposedPort, sshUsername, err := ParseCommandLineArguments()
	if err != nil {
		log.Fatalf("Error while parsing command line arguments: %s", err)
	}

	DisplayNB()

	DisplaySettings(sshUsername, sshServer, exposedPort, targetHost)

	// Initialize SSH listener 0.0.0.0:<expose_port> on the SSH server <ssh_server>:<ssh_port>
	sshListener, err := InitializeSSHListener(sshUsername, sshServer, exposedPort)
	if err != nil {
		log.Fatalf("Error initializing SSH listener %s", err)
	}

	// Vamos muchachos!
	for {
		Forward(sshListener, targetHost)
	}
}
