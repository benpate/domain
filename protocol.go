package domain

import "strings"

// Protocol returns the protocol to use for a given hostname.
// Local domains return `http://`, while all other domains return `https://`
func Protocol(hostname string) string {
	if IsLocalhost(hostname) {
		return "http://"
	}
	return "https://"
}

// HasProtocol returns TRUE if the provided URL includes a protocol string
func HasProtocol(url string) bool {

	// Case-insensitive check for protocol
	url = strings.ToLower(url)

	if strings.HasPrefix(url, "http://") {
		return true
	}

	if strings.HasPrefix(url, "https://") {
		return true
	}

	return false
}

// AddProtocol adds a protocol to a URL if it does not already exist.
// Local domains use http:// and all other domains use https://
func AddProtocol(url string) string {

	if HasProtocol(url) {
		return url
	}

	return Protocol(url) + url
}
