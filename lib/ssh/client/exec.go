package client

import (
	"bytes"
	"errors"
	"fmt"
	"golang.org/x/crypto/ssh"
)

// ProbeBindByPort tries to probe real listen address for the exposed service (SSH forwarder)
func ProbeBindByPort(client ssh.Client, port int) (string, error) {
	stdout, _, err := executeCommand(
		client,
		getBindByPortCommand(port),
	)
	if err != nil {
		return "", err
	}
	if stdout == "" {
		return "", errors.New("No output from probe command :-/")
	}

	return stdout, nil
}

func getBindByPortCommand(port int) string {
	return fmt.Sprintf("netstat -nl | grep \"^tcp.*%d\" | head -n1 | awk '{print $4}' | sed 's/:.*//'", port)
}

func executeCommand(client ssh.Client, command string) (string, string, error) {
	session, err := client.NewSession()
	if err != nil {
		return "", "", err
	}
	defer session.Close()

	var outBuf bytes.Buffer
	var errBuf bytes.Buffer
	session.Stdout = &outBuf
	session.Stderr = &errBuf

	session.Run(command)

	return outBuf.String(), errBuf.String(), nil
}
