package main

import (
	"testing"
)

var prettyJsonTests = []string{
	`{"foo": "bar"}`,
	`["foo", "bar"]`,
}

func TestPrettyJson(t *testing.T) {
	for _, tt := range prettyJsonTests {
		pretty, err := Pretty("application/json", []byte(tt))

		if err != nil {
			t.Errorf("Pretty(%q) err <%s>", tt, err)
		} else {
			if pretty == tt {
				t.Errorf("Pretty(%q) = <%s>", tt, pretty)
			}
		}
	}
}
