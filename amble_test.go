package main

import (
	"testing"
)

type H map[string]string
type E []string

var parseHeaderTests = []struct {
	raw string
	out H
	err E
}{
	{"Host: example.com", H{"Host": "example.com"}, nil},
	{"Host:             example.com", H{"Host": "example.com"}, nil},
	{"Host:example.com", H{"Host": "example.com"}, nil},
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

	if herr, ok := err.(*ParseHeaderError); ok {
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

	return false
}

func TestParseHeaders(t *testing.T) {
	for _, tt := range parseHeaderTests {
		headers, err := ParseHeaders(tt.raw)
		// TODO: why doesn't reflect.DeepEqual work?
		if !HeaderEqual(headers, tt.out) {
			t.Errorf("ParseHeader(%q) = <%s> want <%s>", tt.raw, headers, tt.out)
		}
		if !ErrorEqual(err, tt.err) {
			t.Errorf("ParseHeader(%q) err <%s> want <%s>", tt.raw, err, tt.err)
		}
	}
}
