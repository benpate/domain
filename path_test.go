package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTrimHostname(t *testing.T) {
	require.Equal(t, "/", TrimHostname("https://localhost"))
	require.Equal(t, "/", TrimHostname("https://localhost/"))
	require.Equal(t, "/path", TrimHostname("https://localhost/path"))
	require.Equal(t, "/path?and=query", TrimHostname("https://localhost/path?and=query"))

	require.Equal(t, "/", TrimHostname("http://localhost"))
	require.Equal(t, "/", TrimHostname("http://localhost/"))
	require.Equal(t, "/path", TrimHostname("http://localhost/path"))
	require.Equal(t, "/path?and=query", TrimHostname("http://localhost/path?and=query"))
}
