package main

import (
	"strings"
	"testing"
)

type H map[string]string
type E []string

var parseHeaderTests = []struct {
	raw      string
	out      H
	errLines E
}{
	{"Host: example.com", H{"Host": "example.com"}, nil},
	{"Host:             example.com", H{"Host": "example.com"}, nil},
	{"Host:example.com", H{"Host": "example.com"}, nil},
	{"Host: example.com:123", H{"Host": "example.com:123"}, nil},
	{"a:b\nc:d", H{"a": "b", "c": "d"}, nil},
	{"foo", H{}, E{"foo"}},
	{"foo\na:b", H{"a": "b"}, E{"foo"}},
}

func HeaderEqual(a H, b H) bool {
	if len(a) != len(b) {
		return false
	}

	for k, v := range a {
		if v != b[k] {
			return false
		}
	}

	return true
}

func ErrorEqual(err error, exp E) bool {
	if err == nil && exp == nil {
		return true
	}

	herr, ok := err.(*ParseHeaderError)
	if !ok {
		return false
	}

	// TODO: why doesn't reflect.DeepEqual work?
	if len(herr.FailedLines) != len(exp) {
		return false
	}

	for i, line := range herr.FailedLines {
		if line != exp[i] {
			return false
		}
	}

	return true
}

func TestParseHeaders(t *testing.T) {
	for _, tt := range parseHeaderTests {
		headers, err := ParseHeaders(strings.NewReader(tt.raw))
		// TODO: why doesn't reflect.DeepEqual work?
		if !HeaderEqual(headers, tt.out) {
			t.Errorf("ParseHeader(%q) = <%s> want <%s>", tt.raw, headers, tt.out)
		}
		if !ErrorEqual(err, tt.errLines) {
			t.Errorf("ParseHeader(%q) err <%s> want <%s>", tt.raw, err, tt.errLines)
		}
	}
}
