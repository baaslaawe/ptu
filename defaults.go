package main

import (
	"log"
	"math/rand"
	"os/user"
)

var (
	DefaultSSHServer   = DummySSHServer
	DefaultTargetHost  = "localhost:22"
	DefaultExposedPort = GetDefaultExposedPort()
	DefaultSSHUsername = GetDefaultSSHUsername()
)

const (
	DummySSHServer  = "not.a.real.server"
	BaseExposedPort = 10000
	DefaultSSHPort  = 22
)

func GetDefaultExposedPort() int {
	return BaseExposedPort + rand.Intn(10000)
}

func GetDefaultSSHUsername() string {
	currentUser, err := user.Current()
	if err != nil {
		log.Panicf("Unable to get current user name!")
	}

	return currentUser.Username
}
