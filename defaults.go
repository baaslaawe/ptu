package main

import (
	"log"
	"math/rand"
	"os/user"
)

var (
	DefaultSSHServer   = DummySSHServer
	DefaultSSHUsername = getDefaultSSHUsername()
	DefaultSSHPassword = DummySSHPassword
	DefaultTargetHost  = "localhost:22"
	DefaultExposedHost = "0.0.0.0"
	DefaultExposedPort = getDefaultExposedPort()
)

const (
	DummySSHServer    = "not.a.real.server"
	DummySSHPassword  = ""
	BaseExposedPort   = 10000
	DefaultSSHPort    = 22
	DefaultTargetPort = 80
)

func getDefaultExposedPort() int {
	return BaseExposedPort + rand.Intn(10000)
}

func getDefaultSSHUsername() string {
	currentUser, err := user.Current()
	if err != nil {
		log.Panicf("Unable to get current user name!")
	}

	return currentUser.Username
}
