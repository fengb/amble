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
	Url    string
	Status int
	Header http.Header
	Body   []byte
}

func Fetch(headers map[string]string, url string) (FetchResult, error) {
	result := FetchResult{Url: url}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return result, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()
	result.Status = resp.StatusCode
	result.Header = resp.Header

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	result.Body = body
	return result, nil
}
