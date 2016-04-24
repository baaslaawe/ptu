package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/user"
	"regexp"
	"strconv"
	"time"

	"github.com/ivanilves/gopack/sshtunnel"
)

const (
	version  = "v0.5.0"
	codename = "Amante Latino"

	defaultExposedBind = "0.0.0.0"
	baseExposedPort    = 10000
	retrySeconds       = 5
)

var (
	defaultExposedPort = getDefaultExposedPort()
)

// Config is a container for ptu configuration
type Config struct {
	SSHServer   string `yaml:"s"`
	SSHUsername string `yaml:"u"`
	SSHPassword string `yaml:"p"`
	TargetHost  string `yaml:"t"`
	ExposedBind string `yaml:"b"`
	ExposedPort int    `yaml:"e"`

	BuildID      string
	FailoverPort int
}

// loadDefaults loads default config, either built-in or from default.yaml file (if it exists)
func loadDefaults() (*Config, error) {
	if !doesYAMLExist("default", getYAMLConfigDir()) || isTailored() {
		return getBuiltinDefaults(), nil
	}

	d, err := loadYAML("default", getYAMLConfigDir(), getBuiltinDefaults())
	if err != nil {
		return nil, err
	}

	return d, nil
}

// getBuiltinDefaults tells built-in parameter defaults. NB! Built-in defaults are also used for tailoring!
func getBuiltinDefaults() *Config {
	return &Config{
		SSHServer:   "",
		SSHUsername: getDefaultSSHUsername(),
		SSHPassword: "",
		TargetHost:  "localhost:22",
		ExposedBind: defaultExposedBind,
		ExposedPort: defaultExposedPort,

		BuildID:      "Vanilla",
		FailoverPort: 0,
	}
}

// isTailored tells, if application was tailored
func isTailored() bool {
	return getBuiltinDefaults().BuildID != "Vanilla"
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

func isHelpRequested() bool {
	if len(os.Args) < 2 {
		return false
	}

	return regexp.MustCompile(`^(-h|--help|help)$`).MatchString(os.Args[1])
}

func parseArguments(d *Config) (*Config, error) {
	var s = flag.String("s", d.SSHServer, "SSH server (host[:port]) to connect")
	var u = flag.String("u", d.SSHUsername, "username to connect SSH server")
	var p = flag.String("p", "N/A", "password to authenticate against SSH server")
	var t = flag.String("t", d.TargetHost, "target host:port we will forward connections to")
	var b = flag.String("b", d.ExposedBind, "bind (listener) to expose on the SSH server side")
	var e = flag.Int("e", d.ExposedPort, "port to expose and forward on the SSH server side")

	var yaml = flag.String("c", "", "YAML config name to load from '~/.ptu' directory")

	flag.Parse()

	c := &Config{SSHServer: *s, SSHUsername: *u, SSHPassword: *p, TargetHost: *t, ExposedBind: *b, ExposedPort: *e}

	// We do NOT want to show any password in a help message, so we set it here
	if c.SSHPassword == "N/A" {
		c.SSHPassword = d.SSHPassword
	}

	// Build ID & failover port are always taken from the defaults
	c.BuildID = d.BuildID
	c.FailoverPort = d.FailoverPort

	// If we use YAML config file, we load settings from it first ...
	if *yaml != "" {
		c, err := loadYAML(*yaml, getYAMLConfigDir(), c)
		if err != nil {
			return nil, err
		}

		// ... and override YAML config settings with ones passed by arguments
		flag.Visit(func(f *flag.Flag) {
			switch f.Name {
			case "s":
				c.SSHServer = f.Value.String()
			case "u":
				c.SSHUsername = f.Value.String()

			case "p":
				c.SSHPassword = f.Value.String()

			case "t":
				c.TargetHost = f.Value.String()

			case "b":
				c.ExposedBind = f.Value.String()

			case "e":
				e, _ := strconv.Atoi(f.Value.String())
				c.ExposedPort = e
			}
		})
	}

	// This the only major validation done outside sshtunnel package
	if c.SSHServer == "" {
		return nil, errors.New("SSH server not defined (try to run program with `--help` option)")
	}

	return c, nil
}

func main() {
	if isHelpRequested() {
		fmt.Printf(helpMessage())
		os.Exit(1)
	}

	fmt.Printf(welcomeMessage())

	// Load defaults: built-in or from file (only if it exists)
	d, err := loadDefaults()
	if err != nil {
		log.Fatalf("Unable to load defaults: %s", err)
	}

	// Merge default config with params passed as command line arguments
	c, err := parseArguments(d)
	if err != nil {
		log.Fatalf("Error while parsing command line arguments: %s", err)
	}

	sshServer := c.SSHServer                                         // This variable may be mutated by failover API, if failover is enabled
	failoverServer := failoverSSHServer(c.SSHServer, c.FailoverPort) // If failover is enabled, go here on main server connection failure
	var failoverAPIError error                                       // A placeholder for the failover API error, if OMG we will spot one :/

	for {
		// Initialize instance of the SSH tunnel (at least try to!)
		tunnel, errT := sshtunnel.NewInstance(sshServer, c.SSHUsername, c.SSHPassword, c.TargetHost, c.ExposedBind, c.ExposedPort)
		if errT != nil {
			log.Printf("Error initializing SSH tunnel: %s (will retry)", errT)
			time.Sleep(retrySeconds * time.Second)

			if c.FailoverPort != 0 {
				if sshServer == c.SSHServer {
					failoverAPIError = failoverAPIRequest(sshServer, c.BuildID)
					if failoverAPIError != nil {
						log.Printf("Failover API request failed: %s", failoverAPIError)
					} else {
						sshServer = failoverServer
						log.Printf("* Will try failover server: %s", failoverServer)
					}
				} else {
					sshServer = c.SSHServer
				}
			}

			continue
		} else {
			log.Printf("[OK] SSH tunnel initialized!")
		}

		fmt.Printf(configInfo(*tunnel, c.BuildID)) // Inform about tunnel settings

		// Check, if listener really listens to specified bind address (GatewayPorts thing)
		b, errB := tunnel.ProbeExposedBind()
		if errB != nil {
			log.Printf("WARNING: %s", errB)
		}
		if tunnel.ExposedBind() != b && b != "0.0.0.0" {
			fmt.Printf(gatewayPortsNB())
		}

		// Vamos bandidos!!!
		errF := tunnel.Forward()
		if errF != nil {
			log.Printf("[OMG] Failure on forwarder: %s (will retry)", errF)
			time.Sleep(retrySeconds * time.Second)
		}

	}
}
