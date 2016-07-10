package main

import (
	"os/user"
	"testing"

	"github.com/ivanilves/gopack/sshtunnel"
)

func getSystemUsername() string {
	user, _ := user.Current()

	return user.Username
}

//
// Configurations
//

// ... with agent (safe to use your localhost)
// ... valid
var av = &Config{
	SSHServer:   "127.0.0.1:22",
	SSHUsername: getSystemUsername(),
	ExposedBind: "0.0.0.0",
	ExposedPort: 8888,
}

// ... invalid
var ai = &Config{
	SSHServer:   "127.0.0.1:22",
	SSHUsername: "hijodeputa",
	ExposedBind: "0.0.0.0",
	ExposedPort: 8888,
}

// Try to bring up tunnel instances...
func TestSSHTunnelInstance_with_Agent_and_ValidCredentials(t *testing.T) {
	_, err := sshtunnel.NewInstance(av.SSHServer, av.SSHUsername, av.SSHPassword, av.TargetHost, av.ExposedBind, av.ExposedPort)
	if err != nil {
		t.Error("Unable to connect with agent and valid credentials")
	}
}

func TestSSHTunnelInstance_with_Agent_and_InvalidUsername(t *testing.T) {
	_, err := sshtunnel.NewInstance(ai.SSHServer, ai.SSHUsername, ai.SSHPassword, ai.TargetHost, ai.ExposedBind, ai.ExposedPort)
	if err == nil {
		t.Error("Was able to connect with agent and invalid username :-/")
	}
}

// Try to load `data/fistro.yaml` and examine what we get...
func TestYAMLLoader(t *testing.T) {
	name := "fistro"
	dir := "data"

	c, errL := loadYAML(name, dir, &Config{})
	if errL != nil {
		t.Error(errL)
	}

	if c.SSHServer != "fistro.org:22" {
		t.Error("Unexpected SSH server:", c.SSHServer)
	}

	if c.SSHUsername != "pecador" {
		t.Error("Unexpected SSH username:", c.SSHUsername)
	}

	if c.SSHPassword != "molotov" {
		t.Error("Unexpected SSH password:", c.SSHPassword)
	}

	if c.TargetHost != "pradera:80" {
		t.Error("Unexpected target host:", c.TargetHost)
	}
}
