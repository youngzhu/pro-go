package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func main() {
	go http.ListenAndServe(":5000", nil)
	time.Sleep(time.Second)

	var builder strings.Builder
	err := json.NewEncoder(&builder).Encode(Products[0])
	if err == nil {
		reqURL, err := url.Parse("http://localhost:5000/echo")
		if err == nil {
			req := http.Request{
				Method: http.MethodPost,
				URL:    reqURL,
				Header: map[string][]string{
					"Content-Type": {"application/json"},
				},
				Body: io.NopCloser(strings.NewReader(builder.String())),
			}

			resp, err := http.DefaultClient.Do(&req)
			if err != nil {
				Printfln("Error: %v", err.Error())
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode == http.StatusOK {
				io.Copy(os.Stdout, resp.Body)
			} else {
				Printfln("Status Code: %v", resp.StatusCode)
			}
		}

	}

	if err != nil {
		Printfln("Error: %v", err.Error())
	}

}
