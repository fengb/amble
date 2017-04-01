package main

import (
	"io/ioutil"
	"net/http"
)

var client = &http.Client{
	CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	},
}

type FetchResult struct {
	Header http.Header
	Body   []byte
}

func Fetch(headers map[string]string, url string) (FetchResult, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return FetchResult{}, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return FetchResult{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return FetchResult{}, err
	}

	return FetchResult{Header: resp.Header, Body: body}, err
}
