// Package domain contains some nifty tools for manipulating domains and domain names.
package domain

import (
	"net/url"
	"strings"
)

// Protocol returns the protocol to use for a given hostname.
// Local domains return `http://`, while all other domains return `https://`
func Protocol(hostname string) string {
	if IsLocalhost(hostname) {
		return "http://"
	}
	return "https://"
}

// NameOnly removes the protocol and port from a hostname
func NameOnly(value string) string {

	value = strings.ToLower(value)

	// Remove HTTP/HTTPS protocol, if present
	if strings.HasPrefix(value, ProtocolHTTPS) {
		value = strings.TrimPrefix(value, ProtocolHTTPS)
	} else if strings.HasPrefix(value, ProtocolHTTP) {
		value = strings.TrimPrefix(value, ProtocolHTTP)
	}

	value, _, _ = strings.Cut(value, "/") // Remove path values
	value, _, _ = strings.Cut(value, ":") // Remote port values

	return value
}

func TrimHostname(value string) string {

	// Value must parse as a valid URL, otherwise return empty string
	parsed, err := url.Parse(value)

	if err != nil {
		return ""
	}

	// Force leading slash, if missing
	if parsed.Path == "" {
		parsed.Path = "/"
	}

	// If query is empty, then just return the path
	if parsed.RawQuery == "" {
		return parsed.Path
	}

	// Return path and query string
	return parsed.Path + "?" + parsed.RawQuery
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
