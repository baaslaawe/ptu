package display

import (
	"../config"
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

func PrintConfig(config config.Config) {
	log.Printf("SSH server    : %s", config.SSHServer)
	log.Printf("SSH username  : %s", config.SSHUsername)
	log.Printf("SSH use agent : %v", config.SSHUseAgent)
	log.Printf("Target host   : %s", config.TargetHost)
	log.Printf("--------------------------------------------------------------------------------")
	log.Printf("Connect to (use your specific client software): %s", config.ConnectTo)
	log.Printf("--------------------------------------------------------------------------------")
}
