package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"os/user"
	"strings"
)

// LoadYAML loads config from named YAML file and merges it into default config
func LoadYAML(name string, d *Config) (*Config, error) {
	f := getYAMLFileName(name)

	data, errF := ioutil.ReadFile(f)
	if errF != nil {
		return nil, errF
	}

	errY := yaml.Unmarshal(data, d)
	if errY != nil {
		return nil, errY
	}

	return d, nil
}

// YAMLExists checks, if named YAML config file exists
func YAMLExists(name string) bool {
	y := getYAMLFileName(name)
	if y == "" {
		return false
	}

	f, err := os.Stat(y)
	if err != nil {
		return false
	}
	if f.IsDir() {
		return false
	}

	return true
}

func getYAMLFileName(name string) string {
	h := getUserHomeDir()
	if h == "" {
		return ""
	}

	s := []string{h, "/.ptu/", name, ".yaml"}

	return strings.Join(s, "")
}

func getConfigDir() string {
	h := getUserHomeDir()
	if h == "" {
		return ""
	}

	s := []string{h, ".ptu"}

	return strings.Join(s, "/")
}

func getUserHomeDir() string {
	u, err := user.Current()
	// TODO:
	// Replace this quick workaround with a real solution
	// os/user should work everywhere, reading ENV is bad.
	if err != nil {
		return os.Getenv("HOME")
	}

	return u.HomeDir
}
