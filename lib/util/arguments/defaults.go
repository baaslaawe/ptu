package arguments

import (
	"log"
	"math/rand"
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
		log.Panicf("Unable to get current user name!")
	}

	return currentUser.Username
}
