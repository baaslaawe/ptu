package display

import (
	"fmt"
	"os"
)

func PrintHelpAndExit() {
	fmt.Printf("%s %s\n", name, version)
	fmt.Printf("\n")
	fmt.Printf("Usage: %s -s <ssh_server>[:<ssh_port>] [OPTIONS]\n", name)
	fmt.Printf("--\n")
	fmt.Printf("OPTIONS := -t <target_host>:<target_port> -e <expose_port> -u <ssh_username> -p <ssh_password>\n")

	os.Exit(1)
}

func PrintGatewayPortsNB() {
	fmt.Println("NB!")
	fmt.Println("NB! You may need to enable 'GatewayPorts' option on your SSH server!")
	fmt.Println("NB!")
}

func PrintConfig(
	sshServer string,
	sshUsername string,
	sshUseAgent bool,
	targetHost string,
	connectTo string,
) {
	printSeparator()
	fmt.Println("SSH server    :", sshServer)
	fmt.Println("SSH username  :", sshUsername)
	fmt.Println("SSH use agent :", sshUseAgent)
	fmt.Println("Target host   :", targetHost)
	printSeparator()
	fmt.Println("Connect to (use your specific client software): ", connectTo)
	printSeparator()
}

func printSeparator() {
	fmt.Println("--------------------------------------------------------------------------------")
}
