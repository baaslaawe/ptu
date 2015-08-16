package main

import "regexp"

const (
	TCPPortMIN     = 1
	TCPPortMAX     = 65535
	DefaultSSHPort = 22
)

var (
	HelpArgumentRegexp = regexp.MustCompile(`^(-h|--help)$`)
	HostWithPortRegexp = regexp.MustCompile(`.*:\d+$`)
)
