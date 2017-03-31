package main

import (
	"testing"
)

var fullUrlsTests = []struct {
	headers H
	url     string
	fullUrl string
}{
	{nil, "http://example.com", "http://example.com"},
	{nil, "https://example.com", "https://example.com"},
	{nil, "http://example.com:3000", "http://example.com:3000"},
	{nil, "example.com:3000", "http://example.com:3000"},
	{H{"Origin": "http://example.com"}, "/", "http://example.com/"},
	{H{"Origin": "https://example.com:3000"}, "foobar", "https://example.com:3000/foobar"},
	{H{"Host": "example.com"}, "/", "http://example.com/"},
	{H{"Host": "example.com:3000"}, "foobar", "http://example.com:3000/foobar"},
}

func TestFullUrls(t *testing.T) {
	for _, tt := range fullUrlsTests {
		fullUrls := FullUrls(tt.headers, []string{tt.url})

		if fullUrls[0] != tt.fullUrl {
			t.Errorf("FullUrls(%q, %q) = <%s> want <%s>", tt.headers, tt.url, fullUrls[0], tt.fullUrl)
		}
	}
}
