package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNameOnly(t *testing.T) {
	require.Equal(t, "localhost", NameOnly("localhost"))
	require.Equal(t, "veronica.local", NameOnly("veronica.local"))
	require.Equal(t, "localhost", NameOnly("https://localhost"))
	require.Equal(t, "localhost", NameOnly("https://localhost/"))
	require.Equal(t, "localhost", NameOnly("https://localhost/path"))
	require.Equal(t, "localhost", NameOnly("https://localhost/path?and=query"))

	require.Equal(t, "localhost", NameOnly("http://localhost"))
	require.Equal(t, "localhost", NameOnly("http://localhost/"))
	require.Equal(t, "localhost", NameOnly("http://localhost/path"))
	require.Equal(t, "localhost", NameOnly("http://localhost/path?and=query"))
}
