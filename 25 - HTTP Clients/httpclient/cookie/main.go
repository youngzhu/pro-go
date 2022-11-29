package main

import (
	"io"
	"net/http"
	"net/http/cookiejar"
	"os"
	"time"
)

func main() {
	go http.ListenAndServe(":5000", nil)
	time.Sleep(time.Second)

	// 去掉这块代码，值不会累加
	// 好像一种记忆功能
	Printfln("Jar before: %v", http.DefaultClient.Jar)
	jar, err := cookiejar.New(nil)
	if err == nil {
		http.DefaultClient.Jar = jar
	}
	Printfln("Jar After: %v", http.DefaultClient.Jar)

	for i := 0; i < 3; i++ {
		req, err := http.NewRequest(http.MethodPost,
			"http://localhost:5000/cookie",
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
}
