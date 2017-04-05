package main

import (
	"io"
	"io/ioutil"
	"regexp"
	"strings"
)

var headerSplitter = regexp.MustCompile(": *")

type ParseHeaderError struct {
	FailedLines []string
}

func (e *ParseHeaderError) Error() string {
	return strings.Join(e.FailedLines, ", ")
}

func ParseHeaders(r io.Reader) (map[string]string, error) {
	raw, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	headers := make(map[string]string)
	failedLines := []string{}
	for _, line := range strings.Split(string(raw), "\n") {
		tokens := headerSplitter.Split(line, 2)
		if len(tokens) == 2 {
			headers[tokens[0]] = string(strings.TrimSpace(tokens[1]))
		} else {
			failedLines = append(failedLines, line)
		}
	}

	if len(failedLines) == 0 {
		return headers, nil
	} else {
		return headers, &ParseHeaderError{FailedLines: failedLines}
	}
}
