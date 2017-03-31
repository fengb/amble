package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var client = &http.Client{}

func stream(url string, headers map[string]string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
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
	fmt.Println(url, resp.Status)
	resp.Header.Write(os.Stdout)
	fmt.Println(string(body))
	return nil
}

func isInteractive(f *os.File) bool {
	info, err := f.Stat()
	if err != nil {
		return false
	}

	return (info.Mode() & os.ModeCharDevice) == os.ModeCharDevice
}

func main() {
	headers := make(map[string]string)
	if !isInteractive(os.Stdin) {
		pipeHeaders, err := ParseHeaders(os.Stdin)
		if pipeHeaders == nil {
			panic(err)
		}

		for k, v := range pipeHeaders {
			headers[k] = v
		}
	}
	stream("http://google.com/", headers)
}
