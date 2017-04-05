package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

var defaultHeader = H{"Server": "mock"}

var fetchTests = []struct {
	status  int
	headers H
	body    []byte
}{
	{200, H{}, []byte("Not e")},
	{204, defaultHeader, nil},
	{302, H{"Location": "http://example.com"}, []byte("Redirect")},
	{422, defaultHeader, []byte("why")},
	{500, defaultHeader, []byte("buh")},
}

func mockHandler(w http.ResponseWriter, req *http.Request) {
	for _, data := range fetchTests {
		if req.URL.String() != "/"+strconv.Itoa(data.status) {
			continue
		}

		for k, v := range data.headers {
			w.Header().Set(k, v)
		}
		w.WriteHeader(data.status)

		if data.body != nil {
			w.Write(data.body)
		}

		return
	}

	panic("URL unknown: " + req.URL.String())
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

	server := httptest.NewServer(http.HandlerFunc(mockHandler))
	defer server.Close()

	for _, tt := range fetchTests {
		result, err := Fetch(headers, server.URL+"/"+strconv.Itoa(tt.status))
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

func TestFetchAll(t *testing.T) {
	headers := make(map[string]string)

	server := httptest.NewServer(http.HandlerFunc(mockHandler))
	defer server.Close()

	urls := []string{
		server.URL + "/200",
		server.URL + "/204",
		server.URL + "/302",
	}

	results := []FetchResult{}
	for result := range FetchAll(headers, urls...) {
		results = append(results, result)
	}

	if len(results) != len(urls) {
		t.Errorf("FetchAll() len = <%d> want <%d>", len(results), len(urls))
	}
	for i, result := range results {
		if result.Url != urls[i] {
			t.Errorf("FetchAll() [%d].Url = <%s> want <%s>", i, result.Url, urls[i])
		}
	}
}
