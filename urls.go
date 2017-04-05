package main

import (
	"errors"
	"fmt"
	"path"
	"regexp"
)

var startsWithProtocol = regexp.MustCompile("^[a-z]+://")
var startsWithDomain = regexp.MustCompile("^[a-z]+(\\.[a-z]+|:[0-9]+)")

func FullUrl(headers map[string]string, url string) (string, error) {
	if startsWithProtocol.MatchString(url) {
		return url, nil
	} else if startsWithDomain.MatchString(url) {
		return "http://" + url, nil
	} else if origin, ok := headers["Origin"]; ok {
		return origin + path.Join("/", url), nil
	} else if host, ok := headers["Host"]; ok {
		return "http://" + host + path.Join("/", url), nil
	}
	return "", errors.New("cannot parse")
}

type FullUrlsError struct {
	Headers    map[string]string
	FailedUrls []string
}

func (e *FullUrlsError) Error() string {
	return fmt.Sprintf("Cannot generate urls <%q> with headers <%q>", e.FailedUrls, e.Headers)
}

func FullUrls(headers map[string]string, urls []string) ([]string, error) {
	fullUrls := []string{}
	failedUrls := []string{}
	for _, url := range urls {
		fullUrl, err := FullUrl(headers, url)
		if err == nil {
			fullUrls = append(fullUrls, fullUrl)
		} else {
			failedUrls = append(failedUrls, url)
		}
	}

	if len(failedUrls) == 0 {
		return fullUrls, nil
	} else {
		return fullUrls, &FullUrlsError{Headers: headers, FailedUrls: failedUrls}
	}
}
