package main

import (
	"encoding/json"
	"io"
	"net/http"
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
		req, err := http.NewRequest(http.MethodPost,
			"http://localhost:5000/echo",
			// 好像没区别啊
			//io.NopCloser(strings.NewReader(builder.String())))
			strings.NewReader(builder.String()))
		if err == nil {
			req.Header["Content-Type"] = []string{"application/json"}

			resp, err := http.DefaultClient.Do(req)
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
