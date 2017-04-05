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
	Error  error
}

func Fetch(headers map[string]string, url string) (FetchResult, error) {
	result := FetchResult{Url: url}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		result.Error = err
		return result, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		result.Error = err
		return result, err
	}
	defer resp.Body.Close()
	result.Status = resp.StatusCode
	result.Header = resp.Header

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		result.Error = err
		return result, err
	}

	result.Body = body
	return result, nil
}

func FetchAll(headers map[string]string, urls ...string) <-chan FetchResult {
	c := make(chan FetchResult)
	go func() {
		mailboxes := []chan FetchResult{}
		for _, url := range urls {
			mailbox := make(chan FetchResult)
			go func(url string) {
				result, _ := Fetch(headers, url)
				mailbox <- result
				close(mailbox)
			}(url)
			mailboxes = append(mailboxes, mailbox)
		}

		for _, mailbox := range mailboxes {
			c <- <-mailbox
		}

		close(c)
	}()
	return c
}
