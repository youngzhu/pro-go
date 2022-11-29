package main

import (
	"io"
	"net/http"
	"os"
	"time"
)

// 默认的重定向最多 10 次

func main() {
	go http.ListenAndServe(":5000", nil)
	time.Sleep(time.Second)

	req, err := http.NewRequest(http.MethodPost,
		"http://localhost:5000/redirect1",
		nil)
	if err == nil {
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

	if err != nil {
		Printfln("Error: %v", err.Error())
	}

}
