package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func stream(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	return nil
}

func main() {
	stream("http://google.com/")
}
