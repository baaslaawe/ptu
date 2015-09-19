package config

// Config is a container for ptu configuration
type Config struct {
	SSHServer   string
	SSHUsername string
	SSHPassword string
	SSHUseAgent bool
	TargetHost  string
	ExposedBind string
	ExposedPort int
	ExposedHost string
	ConnectTo   string
}
