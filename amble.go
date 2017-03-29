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
