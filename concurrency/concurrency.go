package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(3)
	request := http.Request{
		Method: "GET",
		URL: &url.URL{
			Host:   "golang.free.beeceptor.com",
			Scheme: "https",
		},
	}
	requestSender := func() {
		response, err := http.Get(request.URL.String())
		if err != nil {
			panic(err)
		}
		defer response.Body.Close()

		responseBody, err := io.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(responseBody), response.StatusCode)

		wg.Done()
	}

	i := 3
	for i > 0 {
		go requestSender()
		i--
	}
	wg.Wait()
}
