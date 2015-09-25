package main

import (
	"./lib/net/forwarder"
	"./lib/ssh/client"
	"./lib/ssh/listener"
	"./lib/util/config"

	"errors"
	"net/http"
	"os/user"
	"testing"
	"time"
)

func getSystemUsername() string {
	user, _ := user.Current()

	return user.Username
}

//
// Configurations
//

// One with agent (safe to use your localhost)
var a = &config.Config{
	SSHServer:   "127.0.0.1:22",
	SSHUsername: getSystemUsername(),
	SSHUseAgent: true,
	ExposedHost: "0.0.0.0:8888",
}

// One with password (SDF public account)
var p = &config.Config{
	SSHServer:   "sdf.org:22",
	SSHUsername: "ptu",
	SSHPassword: "T0mmyTheCatI5MyName",
	SSHUseAgent: false,
	ExposedHost: "127.0.0.1:9999",
}

func TestSSHClient_with_Agent_and_ValidCredentials(t *testing.T) {
	_, err := client.New(a.SSHServer, a.SSHUsername, a.SSHPassword, a.SSHUseAgent)
	if err != nil {
		t.Error("Unable to connect with agent and valid credentials")
	}
}

func TestSSHClient_with_Agent_and_InvalidUsername(t *testing.T) {
	_, err := client.New(a.SSHServer, "hijodeputa", a.SSHPassword, a.SSHUseAgent)
	if err == nil {
		t.Error("Was able to connect with agent and invalid username :-/")
	}
}

func TestSSHClient_with_ValidPassword(t *testing.T) {
	_, err := client.New(p.SSHServer, p.SSHUsername, p.SSHPassword, p.SSHUseAgent)
	if err != nil {
		t.Error("Unable to connect with valid password")
	}
}

func TestSSHClient_with_InvalidPassword(t *testing.T) {
	_, err := client.New(p.SSHServer, p.SSHUsername, "hayquedecirlomas", p.SSHUseAgent)
	if err == nil {
		t.Error("Was able to connect with invalid password :-/")
	}
}

func TestSSHListener_with_ValidExposedHost(t *testing.T) {
	sshClient, _ := client.New(a.SSHServer, a.SSHUsername, a.SSHPassword, a.SSHUseAgent)

	_, err := listener.New(sshClient, a.ExposedHost)
	if err != nil {
		t.Error("Unable to set up listener with valid exposed host")
	}
}

func TestSSHListener_with_InvalidExposedHost(t *testing.T) {
	sshClient, _ := client.New(a.SSHServer, a.SSHUsername, a.SSHPassword, a.SSHUseAgent)

	_, err := listener.New(sshClient, "e8e66ddd26333e68e0cabe5a68c66a16")
	if err == nil {
		t.Error("Was able to set up listener with invalid exposed host :-/")
	}
}

//
// This is the most remarkable testing example:
// It goes all the way from initializing SSH client
// to checking actual forwarding with HTTP request.
//
func TestForwarder(t *testing.T) {
	l := "127.0.0.1:7777"         // listen to this address:port on SSH server side
	h := "info.cern.ch:80"        // our mock forwarding target host:port
	u := "http://127.0.0.1:7777/" // will make outside-in checking HTTP request to this URL

	sshClient, errC := client.New(a.SSHServer, a.SSHUsername, a.SSHPassword, a.SSHUseAgent)
	if errC != nil {
		t.Fatal("[FWD] Unable to initialize client:", errC)
	}

	sshListener, errL := listener.New(sshClient, l)
	if errL != nil {
		t.Fatal("[FWD] Unable to set up listener:", errL)
	}

	// We run HTTP check as asynchronous goroutine
	requestErrors := make(chan error)
	go func() {
		time.Sleep(2 * time.Second)

		resp, errH := http.Get(u)
		if errH != nil {
			requestErrors <- errors.New(errH.Error())
			return
		}
		if resp.StatusCode != 200 {
			requestErrors <- errors.New("Request returned bad status")
			return
		}

		requestErrors <- nil
	}()

	errF := forwarder.Forward(sshListener, h)
	if errF != nil {
		t.Error("Unable to set up connection forwarder:", errF)
	}

	errR := <-requestErrors
	if errR != nil {
		t.Error("Failed to check forwarding via HTTP request:", errR)
	}
}

// Try to load `data/fistro.yaml` and examine what we get...
func TestYAMLLoader(t *testing.T) {
	name := "fistro"
	dir := "data"
	d := &config.Config{}

	c, errL := config.LoadYAML(name, dir, d)
	if errL != nil {
		t.Error(errL)
	}

	c, errV := config.ValidateConfig(c)
	if errV != nil {
		t.Error(errV)
	}

	if c.SSHServer != "fistro.org:22" {
		t.Error("Unexpected SSH server:", c.SSHServer)
	}

	if c.SSHUsername != "pecador" {
		t.Error("Unexpected SSH username:", c.SSHUsername)
	}

	if c.SSHUseAgent != true {
		t.Error("Must use SSH agent")
	}

	if c.TargetHost != "pradera:80" {
		t.Error("Unexpected target host:", c.TargetHost)
	}

	if c.ConnectTo != "fistro.org:11111" {
		t.Error("Unexpected connection host:", c.ConnectTo)
	}
}
