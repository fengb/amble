package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"regexp"
)

var MIME_TYPE_MATCH = regexp.MustCompile("^[a-z]+/[a-z-]+")

func Pretty(contentType string, body []byte) (string, error) {
	mimeType := MIME_TYPE_MATCH.FindString(contentType)
	switch mimeType {
	case "application/json":
		return PrettyJson(body)
	}
	return "", errors.New("no pretty format")
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
