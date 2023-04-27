// Package domain contains some nifty tools for manipulating domains and domain names.
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

// NameOnly removes the protocol and port from a hostname
func NameOnly(host string) string {
	host = strings.TrimPrefix(host, "http://")
	host = strings.TrimPrefix(host, "https://")
	host = strings.Split(host, "/")[0]
	host = strings.Split(host, ":")[0]

	return host
}

// HasProtocol returns TRUE if the provided URL includes a protocol string
func HasProtocol(url string) bool {

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

// IsLocalhost returns TRUE if the hostname is a local domain
func IsLocalhost(hostname string) bool {

	// Nornalize the hostname
	hostname = NameOnly(hostname)
	hostname = strings.ToLower(hostname)

	if hostname == "localhost" {
		return true
	}

	if hostname == "127.0.0.1" {
		return true
	}

	if strings.HasSuffix(hostname, ".local") {
		return true
	}

	if strings.HasPrefix(hostname, "10.") {
		return true
	}

	if strings.HasPrefix(hostname, "192.168") {
		return true
	}

	return false
}
