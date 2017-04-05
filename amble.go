package main

import (
	"fmt"
	"os"
)

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

	fullUrls, err := FullUrls(headers, os.Args[1:])
	if err != nil {
		panic(err)
	}

	for result := range FetchAll(headers, fullUrls...) {
		if result.Error != nil {
			panic(result.Error)
		}
		fmt.Println(result.Url, result.Status)
		result.Header.Write(os.Stdout)
		pretty, err := Pretty(result.Header["Content-Type"][0], result.Body)
		if err == nil {
			fmt.Println(pretty)
		} else {
			fmt.Println(string(result.Body))
		}
	}
}
