package main

import (
	"log"
)

func DisplayHelp(executableName string) {
	log.Printf("Usage: %s -s <ssh_server>[:<ssh_port>] [-t <target_host>:<target_port> -e <expose_port> -u <ssh_username> -p <ssh_password> -f]", executableName)
}

func DisplayNB() {
	log.Print("NB!")
	log.Print("NB! You may need to enable 'GatewayPorts' option on your SSH server!")
	log.Print("NB!")
}

func DisplaySettings(sshUsername string, sshServer string, exposedPort int, targetHost string) {
	log.Printf("SSH username : %s", sshUsername)
	log.Printf("SSH server   : %s", sshServer)
	log.Printf("Exposed port : %d", exposedPort)
	log.Printf("Target host  : %s", targetHost)
}
