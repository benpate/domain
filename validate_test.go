package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidate_Success(t *testing.T) {

	values := []string{
		"localhost",      // localhost is always valid
		"friday.local",   // .local is always valid
		"connor.com",     // simple example name
		"www.connor.com", // three segments
		"you.can.have.many.subdomains.as.you.want.com", // many segments

		// 1-Letter hostnames
		"i.oh1.me",
		"z.com",

		// Punycode Domains
		"xn--stackoverflow.com",
		"stackoverflow.xn--com",
		"xn--fiqa61au8b7zsevnm8ak20mc4a87e.xn--fiqs8s",

		// 2-letter TLDs
		"wow.british-library.uk",
		"stackoverflow.co.uk",
		"domain.com.uk",
		"domain.co.in",
		"domain.uk.edu.in",

		// Numbers and Hyphens
		"stack.com",
		"sta-ck.com",
		"sta---ck.com",
		"9sta--ck.com",
		"sta--ck9.com",
		"stack99.com",
		"99stack.com",
		"sta99ck.com",
		"connor.com:8080", // Ports ignored
	}

	for _, value := range values {
		if result := IsValidHostname(value); !result {
			t.Log(value)
			require.True(t, result)
		}
		require.Equal(t, IsValidHostname(value), !NotValidHostname(value))
	}
}

func TestValidate_Failure(t *testing.T) {

	values := []string{
		"",                   // empty string bad
		"hello.z",            // TLD too short
		"under_score.bad",    // underscores are not allowed
		".leading.dot",       // leading dot is not allowed
		"trailing.dot.",      // trailing dot is not allowed
		"too.many....dots",   // too many dots
		"-leading-dash.bad",  // leading dash is not allowed
		"trailing-dash-.bad", // trailing dash is not allowed
		"this-is-a-very-long-domain-name-that-is-too-long-to-be-valid-because-its-more-than-63-characters.com", // too long
		"this.whole.domain.name.is.too.logn.to.be.valid.because.even.though.each.segment.is.less.than.63.characters.the.whole.thing.is.more.than.253.characters.so.it.should.return.a.false.and.break.anyway.bitches.once.we.actually.put.more.than.253.characters.into.it.my.god.this.is.long", // too long
	}

	for _, value := range values {
		if result := IsValidHostname(value); result {
			t.Log(value)
			require.False(t, result)
		}
	}
}
