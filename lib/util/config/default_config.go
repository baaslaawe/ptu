package config

import (
	"log"
	"math/rand"
	"os"
	"os/user"
	"time"
)

const (
	baseExposedPort   = 10000
	defaultSSHPort    = 22
	defaultTargetPort = 80
)

var (
	defaultSSHUsername = getDefaultSSHUsername()
	defaultExposedBind = "0.0.0.0"
	defaultExposedPort = getDefaultExposedPort()
)

// LoadDefaults() loads default config either built-in or from default.yaml file (if it exists)
func LoadDefaults() (*Config, error) {
	if !YAMLExists("default") {
		return getBuiltinDefaults(), nil
	}

	d, err := LoadYAML("default", getBuiltinDefaults())
	if err != nil {
		return nil, err
	}

	return d, nil
}

func getBuiltinDefaults() *Config {
	return &Config{
		SSHServer:   "",
		SSHUsername: getDefaultSSHUsername(),
		SSHPassword: "",
		SSHUseAgent: true,
		TargetHost:  "localhost:22",
		ExposedBind: defaultExposedBind,
		ExposedPort: defaultExposedPort,
		ExposedHost: joinHostPort(defaultExposedBind, defaultExposedPort),
		ConnectTo:   "",
	}
}

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
