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

	// value 必须是切片
	formData := map[string][]string{
		"name":     {"Kayak "},
		"category": {"Watersprots"},
		"price":    {"279"},
	}

	resp, err := http.PostForm("http://localhost:5000/echo", formData)
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
