package main

import (
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	go http.ListenAndServe(":5000", nil)
	time.Sleep(time.Second)

	resp, err := http.Get("http://localhost:50001/html")
	if err != nil {
		Printfln("Error: %v", err.Error())
		return
	}
	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err == nil {
			defer resp.Body.Close()

			os.Stdout.Write(body)
		}
	} else {
		Printfln("Status Code: %v", resp.StatusCode)
	}
}
