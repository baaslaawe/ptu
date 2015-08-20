package main

type Config struct {
	sshServer           string
	sshUsername         string
	sshPassword         string
	sshUseAgent         bool
	targetHost          string
	exposedHost         string
	exposedPort         int
	exposedBind         string
	failOnNetworkErrors bool
}
