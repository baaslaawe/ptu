package main

import (
	"./lib/ssh/client"
	"./lib/util/arguments"

	"os/user"
	"testing"
)

// A "mock" configuration
var c = &arguments.Config{
	SSHServer:   "127.0.0.1:22",
	SSHUseAgent: true,
	TargetHost:  "www.google.com:443",
	ExposedBind: "0.0.0.0",
	ExposedPort: 9999,
	ExposedHost: "0.0.0.0:9999",
}

func getSystemUsername() string {
	user, _ := user.Current()

	return user.Username
}

func TestSSHClient_WithValidCredentials(t *testing.T) {
	c.SSHUsername = getSystemUsername()

	_, err := client.New(c.SSHServer, c.SSHUsername, c.SSHPassword, c.SSHUseAgent)
	//
	// We should be able to connect to the testing SSH server with valid credentials
	//
	if err != nil {
		t.Error("Unable to connect to the SSH server with valid credentials")
	}
}

func TestSSHClient_WithInvalidUsername(t *testing.T) {
	c.SSHUsername = "hijodeputa"

	_, err := client.New(c.SSHServer, c.SSHUsername, c.SSHPassword, c.SSHUseAgent)
	//
	// We should NOT be able to connect anywhere with this dummy username
	//
	if err == nil {
		t.Error("Was able to connect to the SSH server with invalid username :-/")
	}
}

func TestSSHClient_WithInvalidPassword(t *testing.T) {
	c.SSHUsername = getSystemUsername()
	c.SSHPassword = "hayquedecirlomas"
	c.SSHUseAgent = false

	_, err := client.New(c.SSHServer, c.SSHUsername, c.SSHPassword, c.SSHUseAgent)
	//
	// We should NOT be able to connect anywhere with this dummy password
	//
	if err == nil {
		t.Error("Was able to connect to the SSH server with invalid password :-/")
	}

}
