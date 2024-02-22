package domain

import (
	"regexp"
	"strings"

	"net/netip"
)

var hostnameSegment *regexp.Regexp

func init() {
	hostnameSegment = regexp.MustCompile(`^[a-z0-9-]+$`)
}

func IsValidHostname(hostname string) bool {

	// RULE: Empty string not allowed
	if len(hostname) == 0 {
		return false
	}

	// Remove Ports and Protocols
	hostname = NameOnly(hostname)

	// If the hostname is a valid IP address, then it's valid
	if _, err := netip.ParseAddr(hostname); err == nil {
		return true
	}

	// RULE: Maximum length of 253 characters
	// https://www.rfc-editor.org/rfc/rfc1035
	if len(hostname) > 253 {
		return false
	}

	// Lowercase the hostname for comparisons
	hostname = strings.ToLower(hostname)

	// RULE: Hostname cannot begin or end with a period
	if strings.HasPrefix(hostname, ".") || strings.HasSuffix(hostname, ".") {
		return false
	}

	// Inspect each segment of the hostname
	segments := strings.Split(hostname, ".")
	segmentCount := len(segments)

	for index, segment := range segments {

		segmentLength := len(segment)

		// RULE: Segments must have at least one character
		if segmentLength == 0 {
			return false
		}

		// RULE: TLDs (final segment) must be at least two characters
		if (segmentLength < 2) && (index+1 == segmentCount) {
			return false
		}

		// RULE: Segments cannot be longer than 63 characters
		if segmentLength > 63 {
			return false
		}

		// RULE: Segments cannot begin or end with a dash
		if strings.HasPrefix(segment, "-") || strings.HasSuffix(segment, "-") {
			return false
		}

		// RULE: Each segment must contain only letters, numbers, and dashes
		if !hostnameSegment.MatchString(segment) {
			return false
		}
	}

	return true
}

func NotValidHostname(hostname string) bool {
	return !IsValidHostname(hostname)
}
