package main

import (
	"path"
	"regexp"
	"strings"
)

var startsWithProtocol = regexp.MustCompile("^[a-z]+://")
var startsWithDomain = regexp.MustCompile("^[a-z]+(\\.[a-z]+|:[0-9]+)")

func FullUrl(headers map[string]string, url string) string {
	if startsWithProtocol.MatchString(url) {
		return url
	} else if startsWithDomain.MatchString(url) {
		return "http://" + url
	} else if origin, ok := headers["Origin"]; ok {
		return origin + path.Join("/", url)
	} else if host, ok := headers["Host"]; ok {
		return "http://" + host + path.Join("/", url)
	}
	return ""
}

type FullUrlsError struct {
	FailedUrls []string
}

func (e *FullUrlsError) Error() string {
	return strings.Join(e.FailedUrls, ", ")
}

func FullUrls(headers map[string]string, urls []string) ([]string, *FullUrlsError) {
	fullUrls := []string{}
	failedUrls := []string{}
	for _, url := range urls {
		fullUrl := FullUrl(headers, url)
		if fullUrl != "" {
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
