package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"regexp"
)

var mimeTypeMatch = regexp.MustCompile("^[a-z]+/[a-z-]+")

func Pretty(contentType string, body []byte) (string, error) {
	if len(contentType) == 0 {
		contentType = http.DetectContentType(body)
	}

	mimeType := mimeTypeMatch.FindString(contentType)

	switch mimeType {
	case "application/json":
		return PrettyJson(body)
	}
	return string(body), nil
}

func PrettyJson(body []byte) (string, error) {
	var pretty bytes.Buffer
	err := json.Indent(&pretty, body, "", "  ")
	if err == nil {
		return pretty.String(), nil
	} else {
		return "", err
	}
}
