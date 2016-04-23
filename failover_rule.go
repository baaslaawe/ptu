package main

import "errors"
import "net/http"
import "net"
import "strings"
import "strconv"
import "bytes"

// failoverAPIRequest requests a failover rule creation from API [CloudPort] server
func failoverAPIRequest(sshServer string, failoverBuildId string) error {
	const apiScheme = "http://"
	const apiPath = "/failover_rules"

	apiServer, _, _ := net.SplitHostPort(sshServer)
	apiURL := strings.Join([]string{apiScheme, apiServer, apiPath, "/?id=", failoverBuildId}, "")

	r, err := http.Post(apiURL, "text/plain", bytes.NewBuffer(nil))

	if err != nil {
		return err
	}

	if r.StatusCode != 200 {
		return errors.New(r.Status)
	}

	return nil
}

// failoverSSHServer returns an address of the failover SSH server
func failoverSSHServer(sshServer string, failoverPort int) string {
	host, _, _ := net.SplitHostPort(sshServer)

	return net.JoinHostPort(host, strconv.Itoa(failoverPort))
}
