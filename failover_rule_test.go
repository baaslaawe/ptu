package main

import "testing"
import "errors"

const failoverServer = "staging.cloudport.xyz:11637"
const failoverPort = 443
const failoverBuildID = "965f82"

func TestFailoverAPIRequest(t *testing.T) {
	err := failoverAPIRequest(failoverServer, failoverBuildID)

	if err != nil {
		t.Error(err)
	}
}

func TestFailoverAPIRequest_withInvalidBuildID(t *testing.T) {
	err := failoverAPIRequest(failoverServer, "cabron")

	if err == nil {
		t.Error(errors.New("Reported success with invalid build ID"))
	}
}

func TestFailoverAPIRequest_withInvalidServer(t *testing.T) {
	err := failoverAPIRequest("culo", failoverBuildID)

	if err == nil {
		t.Error(errors.New("Reported success with invalid server"))
	}
}

func TestFailoverSSHServer(t *testing.T) {
	failoverServer := failoverSSHServer(failoverServer, failoverPort)

	if failoverServer != "staging.cloudport.xyz:443" {
		t.Error("Failover server address returned incorrectly")
	}
}
