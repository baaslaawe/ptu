package display

import (
	"fmt"
	"net"
	"os"
)

// PrintHelpAndExit prints help message and exits with code 1
func PrintHelpAndExit() {
	fmt.Printf("%s %s\n", name, version)
	fmt.Println("--")
	fmt.Printf("Usage: %s -s <ssh_server>[:<ssh_port>] [OPTIONS]\n", os.Args[0])
	fmt.Println("")
	fmt.Println("OPTIONS:")
	fmt.Println("  { -u <ssh_username> | -p <ssh_password> }")
	fmt.Println("  { -t <target_host>:<target_port> | -b <exposed_bind> | -e <exposed_port> }")
	fmt.Println("  { -c <config_name> }")
	fmt.Println("")
	fmt.Println("See also: https://github.com/ivanilves/ptu/blob/master/README.md")

	os.Exit(1)
}

// PrintGatewayPortsNB prints warning about SSH server 'GatewayPorts' option
// Read this: http://www.snailbook.com/faq/gatewayports.auto.html
func PrintGatewayPortsNB() {
	printSeparator()
	fmt.Println("NB!")
	fmt.Println("NB! You may need to enable 'GatewayPorts' option on your SSH server!")
	fmt.Println("NB!")
	printSeparator()
}

// PrintConfig prints runtime configuration info
func PrintConfig(
	sshServer string,
	sshUsername string,
	sshUseAgent bool,
	targetHost string,
	exposedPort int,
	connectTo string,
	buildID string,
) {
	h, _, _ := net.SplitHostPort(connectTo)

	if buildID != "" {
		fmt.Printf("Build ID: %s\n", buildID)
		printSeparator()
	}
	fmt.Println("SSH server    :", sshServer)
	fmt.Println("SSH username  :", sshUsername)
	fmt.Println("SSH use agent :", sshUseAgent)
	fmt.Println("Target host   :", targetHost)
	printSeparator()
	fmt.Println("Connect to (use your specific client software):", connectTo)
	fmt.Println("")
	fmt.Println("[ Examples ]")
	fmt.Printf("%7s : ssh -p %d %s\n", "SSH", exposedPort, h)
	fmt.Printf("%7s : curl http://%s/\n", "HTTP", connectTo)
	fmt.Printf("%7s : curl -k https://%s/\n", "HTTPS", connectTo)
	printSeparator()
}

func printSeparator() {
	fmt.Println("--------------------------------------------------------------------------------")
}
