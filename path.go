package domain

import "net/url"

// PathAndQuery removes the protocol, hostname, and port from a URL,
// returning only the path and querystring.
func PathAndQuery(value string) string {

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

// TrimHostname is an alias for PathAndQuery
// deprecated: Use PathAndQuery instead
func TrimHostname(url string) string {
	return PathAndQuery(url)
}
