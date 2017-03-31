package main

import (
	"path"
	"regexp"
)

var startsWithProtocol = regexp.MustCompile("^[a-z]+://")
var startsWithDomain = regexp.MustCompile("^[a-z]+\\.[a-z]+")

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

func FullUrls(headers map[string]string, urls []string) []string {
	fullUrls := []string{}
	for _, url := range urls {
		fullUrl := FullUrl(headers, url)
		fullUrls = append(fullUrls, fullUrl)
	}
	return fullUrls
}
