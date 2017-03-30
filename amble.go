package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var HEADER_SPLITTER = regexp.MustCompile(": *")

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
	failed_lines := []string{}
	for _, line := range strings.Split(string(raw), "\n") {
		tokens := HEADER_SPLITTER.Split(line, 2)
		if len(tokens) == 2 {
			headers[tokens[0]] = string(strings.TrimSpace(tokens[1]))
		} else {
			failed_lines = append(failed_lines, line)
		}
	}

	if len(failed_lines) == 0 {
		return headers, nil
	} else {
		return headers, &ParseHeaderError{FailedLines: failed_lines}
	}
}

var client = &http.Client{}

func stream(url string, headers map[string]string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("%s - %s\n", url, resp.Status)
	resp.Header.Write(os.Stdout)
	fmt.Println(string(body))
	return nil
}

func main() {
	stream("http://google.com/", make(map[string]string))
}
