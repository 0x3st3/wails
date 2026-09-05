package application

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
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

// uriMatchesProxy reports whether uri addresses the given proxy host and port.
// WebView2 raises BasicAuthenticationRequested for origin and proxy challenges
// alike and provides no auth-type field; for proxy challenges the event's Uri is
// the proxy server's own URI, so it is the only reliable discriminator.
func uriMatchesProxy(uri, proxyHost string, proxyPort int) bool {
	if uri == "" || proxyHost == "" {
		return false
	}
	u, err := url.Parse(uri)
	if err != nil {
		return false
	}
	if !strings.EqualFold(u.Hostname(), proxyHost) {
		return false
	}
	port := u.Port()
	if port == "" {
		switch u.Scheme {
		case "http":
			port = "80"
		case "https":
			port = "443"
		default:
			return false
		}
	}
	return port == strconv.Itoa(proxyPort)
}
