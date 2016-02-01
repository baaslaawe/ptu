package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ivanilves/gopack/sshtunnel"
)

// helpMessage returns a simple help message
func helpMessage() string {
	return strings.Join([]string{
		fmt.Sprintf("ptu %s\n", version),
		fmt.Sprintln("--"),
		fmt.Sprintf("Usage: %s -s <ssh_server>[:<ssh_port>] [OPTIONS]\n", os.Args[0]),
		fmt.Sprintln(""),
		fmt.Sprintln("OPTIONS:"),
		fmt.Sprintln("  { -u <ssh_username> | -p <ssh_password> }"),
		fmt.Sprintln("  { -t <target_host>:<target_port> | -b <exposed_bind> | -e <exposed_port> }"),
		fmt.Sprintln("  { -c <config_name> }"),
		fmt.Sprintln(""),
		fmt.Sprintln("See also: https://github.com/ivanilves/ptu/blob/master/README.md"),
	},
		"",
	)
}

// welcomeMessage welcomes you, guapo!
func welcomeMessage() string {
	return strings.Join([]string{
		separator(),
		fmt.Sprintf("> ptu %s \"%s\"\n", version, codename),
		separator(),
	},
		"",
	)
}

// gatewayPortsNB returns a warning message about SSH server 'GatewayPorts' option [not being enabled]
// Read this: http://www.snailbook.com/faq/gatewayports.auto.html
func gatewayPortsNB() string {
	return strings.Join([]string{
		separator(),
		fmt.Sprintln("NB!"),
		fmt.Sprintln("NB! You may need to enable 'GatewayPorts' option on your SSH server!"),
		fmt.Sprintln("NB!"),
		separator(),
	},
		"",
	)
}

// configInfo returns a message with runtime configuration info
func configInfo(tunnel sshtunnel.Instance, buildID string) string {
	return strings.Join([]string{
		separator(),
		fmt.Sprintln("Build ID :", buildID),
		separator(),
		fmt.Sprintln("SSH server    :", tunnel.SSHServer(), "(do NOT connect here, please!)"),
		fmt.Sprintln("SSH username  :", tunnel.SSHUsername()),
		fmt.Sprintln("SSH use agent :", tunnel.SSHUseAgent()),
		fmt.Sprintln("Target host   :", tunnel.TargetHost()),
		separator(),
		fmt.Sprintln(""),
		fmt.Sprintln("CONNECT HERE:", tunnel.ConnectTo()),
		fmt.Sprintln(""),
		fmt.Sprintln("[ Examples ]"),
		fmt.Sprintf("%7s : ssh -p %d %s\n", "SSH", tunnel.ExposedPort(), tunnel.ConnectToAddr()),
		fmt.Sprintf("%7s : sftp -oPort=%d %s\n", "SFTP", tunnel.ExposedPort(), tunnel.ConnectToAddr()),
		fmt.Sprintf("%7s : rdesktop %s\n", "RDP", tunnel.ConnectTo()),
		fmt.Sprintf("%7s : curl http://%s/\n", "HTTP", tunnel.ConnectTo()),
		fmt.Sprintf("%7s : curl -k https://%s/\n", "HTTPS", tunnel.ConnectTo()),
		separator(),
	},
		"",
	)
}

func separator() string {
	return fmt.Sprintln(strings.Repeat("-", 80))
}
