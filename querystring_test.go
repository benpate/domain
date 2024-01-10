package domain

import (
	"net/url"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
)

func TestQueryString(t *testing.T) {

	uri, err := url.ParseQuery("http://localhost:8080/api/v1/employees?name=John&age=30&city=New%20York&city=San%20Diego&city=Los%20Angeles")

	require.Nil(t, err)
	require.Equal(t, "New York", uri.Get("city"))

	require.Equal(t, []string{"New York", "San Diego", "Los Angeles"}, uri["city"])
	spew.Dump(uri)
}
