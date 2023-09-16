package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsLocalhost(t *testing.T) {

	require.True(t, IsLocalhost("localhost"))
	require.False(t, NotLocalhost("localhost"))

	require.True(t, IsLocalhost("127.0.0.1/john"))
	require.True(t, IsLocalhost("10.0.0.4/@john"))
	require.True(t, IsLocalhost("http://localhost:8080/john"))
	require.True(t, IsLocalhost("http://192.168.69.69"))
	require.True(t, IsLocalhost("https://server.local"))

	require.False(t, IsLocalhost("connor.com"))
	require.True(t, NotLocalhost("connor.com"))

	require.False(t, IsLocalhost("http://connor.com"))
	require.False(t, IsLocalhost("https://connor.com/@john"))
	require.False(t, IsLocalhost("https://connor.com:1234/@john"))
}

func TestAddProtocol(t *testing.T) {
	require.Equal(t, "http://localhost/@john", AddProtocol("localhost/@john"))
	require.Equal(t, "http://localhost/@john", AddProtocol("http://localhost/@john"))

	require.Equal(t, "https://connor.com/@john", AddProtocol("connor.com/@john"))
	require.Equal(t, "https://connor.com/@john", AddProtocol("https://connor.com/@john"))
}
