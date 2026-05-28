package application

import (
	"fmt"
	"net/url"
	"strconv"
)

func parseWebviewProxyServer(server string) (host string, port int, scheme string, err error) {
	u, err := url.Parse(server)
	if err != nil {
		return "", 0, "", err
	}
	scheme = u.Scheme
	if scheme != "http" && scheme != "https" {
		return "", 0, "", fmt.Errorf("unsupported proxy scheme %q (want http or https)", scheme)
	}
	host = u.Hostname()
	portStr := u.Port()
	if host == "" || portStr == "" {
		return "", 0, "", fmt.Errorf("proxy URL must include host and port")
	}
	port, err = strconv.Atoi(portStr)
	if err != nil {
		return "", 0, "", fmt.Errorf("invalid proxy port %q: %w", portStr, err)
	}
	return host, port, scheme, nil
}

// webviewProxyServerArg returns the credential-stripped "scheme://host:port"
// string for use as a Chromium --proxy-server argument.
func webviewProxyServerArg(server string) (string, error) {
	host, port, scheme, err := parseWebviewProxyServer(server)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s://%s:%d", scheme, host, port), nil
}
