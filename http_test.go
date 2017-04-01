package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

var defaultHeader = H{"Server": "mock"}

var fetchTests = []struct {
	status  int
	headers H
	body    []byte
}{
	{200, H{}, nil},
	{200, defaultHeader, []byte("Not e")},
	{204, defaultHeader, nil},
	{302, H{"Location": "http://example.com"}, nil},
	{422, defaultHeader, []byte("why")},
	{500, defaultHeader, []byte("buh")},
}

func EqualHeaders(actual http.Header, expected H) bool {
	for k, v := range expected {
		a := actual[k]
		if len(a) != 1 || a[0] != v {
			return false
		}
	}
	return true
}

func TestFetch(t *testing.T) {
	headers := make(map[string]string)
	for _, tt := range fetchTests {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			for k, v := range tt.headers {
				w.Header().Set(k, v)
			}
			w.WriteHeader(tt.status)

			if tt.body != nil {
				w.Write(tt.body)
			}
		}))

		defer server.Close()

		result, err := Fetch(headers, server.URL)
		if err != nil {
			t.Errorf("Fetch(%d) err <%s>", tt.status, err)
		}
		if result.Status != tt.status {
			t.Errorf("Fetch(%d) Status = <%d>", tt.status, result.Status)
		}
		if !EqualHeaders(result.Header, tt.headers) {
			t.Errorf("Fetch(%d) Header = <%s> want <%s>", tt.status, result.Header, tt.headers)
		}
		if !bytes.Equal(result.Body, tt.body) {
			t.Errorf("Fetch(%d) Body = <%s> want <%s>", tt.status, result.Body, tt.body)
		}
	}
}
