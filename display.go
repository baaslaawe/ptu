package main

import (
	"log"
)

func DisplayHelp() {
	log.Printf("%s %s", NAME, VERSION)
	log.Printf("--")
	log.Printf("Usage: %s -s <ssh_server>[:<ssh_port>] [-t <target_host>:<target_port> -e <expose_port> -u <ssh_username> -p <ssh_password> -f]", NAME)
}

func DisplayNB() {
	log.Print("NB!")
	log.Print("NB! You may need to enable 'GatewayPorts' option on your SSH server!")
	log.Print("NB!")
}

func DisplayConfig(config Config) {
	log.Printf("SSH server             : %s", config.sshServer)
	log.Printf("SSH username           : %s", config.sshUsername)
	log.Printf("SSH use agent          : %s", config.sshUseAgent)
	log.Printf("Exposed bind           : %s", config.exposedBind)
	log.Printf("Target host            : %s", config.targetHost)
	log.Printf("Fail on network errors : %s", config.failOnNetworkErrors)
}
