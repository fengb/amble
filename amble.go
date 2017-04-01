package main

import (
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

	for _, fullUrl := range fullUrls {
		Fetch(headers, fullUrl)
	}
}
