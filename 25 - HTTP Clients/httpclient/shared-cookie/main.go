package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"os"
	"time"
)

// 多个client，共享一个cookie

func main() {
	go http.ListenAndServe(":5000", nil)
	time.Sleep(time.Second)

	clients := make([]http.Client, 3)
	jar, err := cookiejar.New(nil)
	for index, client := range clients {
		if err == nil {
			client.Jar = jar
		}

		for i := 0; i < 3; i++ {
			req, err := http.NewRequest(http.MethodPost,
				"http://localhost:5000/cookie",
				nil)
			if err == nil {
				resp, err := client.Do(req)
				if err != nil {
					Printfln("Error: %v", err.Error())
					return
				}

				if resp.StatusCode == http.StatusOK {
					defer resp.Body.Close()

					fmt.Fprintf(os.Stdout, "Client-%v: ", index)
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

}
