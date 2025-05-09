package domain

import (
	"bufio"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHostname(t *testing.T) {

	do := func(request string, expected string) {

		// Read the string as an HTTP request
		reader := io.NopCloser(strings.NewReader(request))
		bufferedReader := bufio.NewReader(reader)
		req, err := http.ReadRequest(bufferedReader)

		// Test results
		require.NoError(t, err)
		require.Equal(t, expected, Hostname(req))
	}

	// Validate localhost
	do(`GET /my/path HTTP/1.1
Host: localhost
User-Agent: RapidAPI/4.2.8 (Macintosh; OS X/15.1.1)

`, "localhost")

	// Validate localhost with a port
	do(`GET /my/path HTTP/1.1
Host: localhost:8080
Connection: close
User-Agent: RapidAPI/4.2.8 (Macintosh; OS X/15.1.1) GCDHTTPRequest

`, "localhost:8080")

	// Validate public address
	do(`GET /my/path HTTP/1.1
Host: awesome.com
User-Agent: RapidAPI/4.2.8 (Macintosh; OS X/15.1.1)

`, "awesome.com")

	// Validate public address with a port
	do(`GET /my/path HTTP/1.1
Host: awesome.com:8080
User-Agent: RapidAPI/4.2.8 (Macintosh; OS X/15.1.1)

`, "awesome.com:8080")

	// Validate public address via X-Forwarded-Host proxy
	do(`GET /my/path HTTP/1.1
Host: localhost
X-Forwarded-Host: awesome.com
User-Agent: RapidAPI/4.2.8 (Macintosh; OS X/15.1.1)

`, "awesome.com")

}
