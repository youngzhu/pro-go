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
		resp, err := http.Post("http://localhost:5000/echo",
			"application/json",
			strings.NewReader(builder.String()))
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
	} else {
		Printfln("Error: %v", err.Error())
	}

}
