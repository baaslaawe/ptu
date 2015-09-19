package config

import (
	"log"
	"math/rand"
	"os"
	"os/user"
	"time"
)

var (
	defaultSSHServer   = dummySSHServer
	defaultSSHUsername = getDefaultSSHUsername()
	defaultSSHPassword = dummySSHPassword
	defaultTargetHost  = "localhost:22"
	defaultExposedBind = "0.0.0.0"
	defaultExposedPort = getDefaultExposedPort()
)

const (
	dummySSHServer    = "some.ssh.server"
	dummySSHPassword  = ""
	baseExposedPort   = 10000
	defaultSSHPort    = 22
	defaultTargetPort = 80
)

func getDefaultExposedPort() int {
	rand.Seed(time.Now().UnixNano())

	return baseExposedPort + rand.Intn(10000)
}

func getDefaultSSHUsername() string {
	currentUser, err := user.Current()
	if err != nil {
		// TODO:
		// Replace this quick workaround with a real solution
		// os/user should work everywhere, reading ENV is bad.
		// Vladimir Titov spotted this bug. Thank you Vladimir!
		if os.Getenv("USER") != "" {
			return os.Getenv("USER")
		}

		log.Printf("Unable to get current user name!")

		return "ptu"
	}

	return currentUser.Username
}
