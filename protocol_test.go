package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddProtocol(t *testing.T) {
	require.Equal(t, "http://localhost/@john", AddProtocol("localhost/@john"))
	require.Equal(t, "http://localhost/@john", AddProtocol("http://localhost/@john"))

	require.Equal(t, "https://connor.com/@john", AddProtocol("connor.com/@john"))
	require.Equal(t, "https://connor.com/@john", AddProtocol("https://connor.com/@john"))
}
