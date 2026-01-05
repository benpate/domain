// Package domain contains some nifty tools for manipulating domains and domain names.
package domain

import (
	"strings"
)

// IsLocalhost returns TRUE if the hostname is a local domain
func IsLocalhost(hostname string) bool {

	// Private networks are defined by RFC 1918
	// https://en.wikipedia.org/wiki/Private_network

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

	// 24-bit block
	if strings.HasPrefix(hostname, "10.") {
		return true
	}

	// 20-bit block
	if strings.HasPrefix(hostname, "172.16") {
		return true
	}

	// 16-bit block
	if strings.HasPrefix(hostname, "192.168") {
		return true
	}

	// TODO: IPv6 private networks

	return false
}

// NotLocalhost returns TRUE if the hostname is NOT a local domain
func NotLocalhost(hostname string) bool {
	return !IsLocalhost(hostname)
}
