package domain

import (
	"strings"
)

// Hostname returns ONLY the hostname, removing
// the protocol, port, path, and querystring from a hostname
func Hostname(value string) string {

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

// NameOnly removes the protocol and port from a hostname
// deprecated: Use Hostname instead
func NameOnly(value string) string {
	return Hostname(value)
}
