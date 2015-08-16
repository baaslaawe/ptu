package main

import (
	"log"
)

func DisplayHelp(executableName string) {
	log.Printf("Usage: %s [-u <ssh_username>] -s <ssh_server>[:<ssh_port>] -p <forward_port> -r <remote_host>[:<remote_port>]", executableName)
}

func DisplayNB() {
	log.Print("NB!")
	log.Print("NB! You may need to enable 'GatewayPorts' option on your SSH server!")
	log.Print("NB!")
}

func DisplaySettings(sshUsername string, sshServer string, forwardPort int, remoteHost string) {
	log.Printf("SSH username : %s", sshUsername)
	log.Printf("SSH server   : %s", sshServer)
	log.Printf("Forward port : %d", forwardPort)
	log.Printf("Remote host  : %s", remoteHost)
}
