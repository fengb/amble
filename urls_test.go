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
	{nil, "example:3000", "http://example:3000"},
	{H{"Origin": "http://example.com"}, "/", "http://example.com/"},
	{H{"Origin": "https://example.com:3000"}, "foobar", "https://example.com:3000/foobar"},
	{H{"Host": "example.com"}, "/", "http://example.com/"},
	{H{"Host": "example.com:3000"}, "foobar", "http://example.com:3000/foobar"},
	{nil, "/", ""},
	{nil, "foobar", ""},
}

func TestFullUrls(t *testing.T) {
	for _, tt := range fullUrlsTests {
		fullUrls, err := FullUrls(tt.headers, []string{tt.url})

		if tt.fullUrl != "" {
			if len(fullUrls) == 0 || fullUrls[0] != tt.fullUrl {
				t.Errorf("FullUrls(%q, %q) = <%s> want <[%s]>", tt.headers, tt.url, fullUrls, tt.fullUrl)
			}
		} else {
			if len(fullUrls) != 0 {
				t.Errorf("FullUrls(%q, %q) = <%s> should err", tt.headers, tt.url, fullUrls)
			}

			if err != nil {
				if uerr, ok := err.(*FullUrlsError); ok {
					if uerr.FailedUrls[0] != tt.url {
						t.Errorf("FullUrls(%q, %q) err <%s> want <%s>", tt.headers, tt.url, err, tt.url)
					}
				} else {
					t.Errorf("FullUrls(%q, %q) unknown err <%s>", tt.headers, tt.url, err)
				}
			}
		}
	}
}
