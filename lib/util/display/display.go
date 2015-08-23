package display

import (
	"log"
	"os"
)

func PrintHelpAndExit() {
	log.Printf("%s %s", name, version)
	log.Printf("--")
	log.Printf("Usage: %s -s <ssh_server>[:<ssh_port>] [OPTIONS]", name)
	log.Printf("--")
	log.Printf("OPTIONS := -t <target_host>:<target_port> -e <expose_port> -u <ssh_username> -p <ssh_password>")

	os.Exit(1)
}

func PrintGatewayPortsNB() {
	log.Print("NB!")
	log.Print("NB! You may need to enable 'GatewayPorts' option on your SSH server!")
	log.Print("NB!")
}

func PrintConfig(
	sshServer string,
	sshUsername string,
	sshUseAgent bool,
	targetHost string,
	connectTo string,
) {
	log.Printf("SSH server    : %s", sshServer)
	log.Printf("SSH username  : %s", sshUsername)
	log.Printf("SSH use agent : %v", sshUseAgent)
	log.Printf("Target host   : %s", targetHost)
	log.Printf("--------------------------------------------------------------------------------")
	log.Printf("Connect to (use your specific client software): %s", connectTo)
	log.Printf("--------------------------------------------------------------------------------")
}
