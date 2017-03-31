package main

import (
	"errors"
	"path"
	"regexp"
	"strings"
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
	FailedUrls []string
}

func (e *FullUrlsError) Error() string {
	return strings.Join(e.FailedUrls, ", ")
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
		return fullUrls, &FullUrlsError{FailedUrls: failedUrls}
	}
}
