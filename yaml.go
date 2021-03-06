package main

import (
	"io/ioutil"
	"os"
	"os/user"
	"strings"

	"gopkg.in/yaml.v2"
)

// loadYAML loads config from named YAML file and merges it into default config
func loadYAML(name string, dir string, d *Config) (*Config, error) {
	f := getYAMLFileName(name, dir)

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

// doesYAMLExist checks, if named YAML config file exists in the specified directory
func doesYAMLExist(name string, dir string) bool {
	y := getYAMLFileName(name, dir)
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

func getYAMLConfigDir() string {
	h := getUserHomeDir()
	if h == "" {
		return ""
	}

	s := []string{h, ".ptu"}

	return strings.Join(s, "/")
}

func getYAMLFileName(name string, dir string) string {
	s := []string{dir, "/", name, ".yaml"}

	return strings.Join(s, "")
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
