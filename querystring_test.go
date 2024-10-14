package domain

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueryString(t *testing.T) {

	uri, err := url.ParseQuery("name=John&age=30&city=New%20York&city=San%20Diego&city=Los%20Angeles")

	require.Nil(t, err)
	require.Equal(t, "John", uri.Get("name"))
	require.Equal(t, "30", uri.Get("age"))
	require.Equal(t, "New York", uri.Get("city"))
	require.Equal(t, []string{"New York", "San Diego", "Los Angeles"}, uri["city"])
}
