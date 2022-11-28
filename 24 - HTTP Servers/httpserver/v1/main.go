package main

import (
	"fmt"
	"io"
	"net/http"
)

type StringHandler struct {
	message string
}

func (sh StringHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	Printfln("Method: %v", r.Method)
	Printfln("URL: %v", r.URL)
	Printfln("HTTP Version: %v", r.Proto)
	Printfln("Host: %v", r.Host)
	for name, val := range r.Header {
		Printfln("Header: %v, Value: %v", name, val)
	}
	Printfln("---")
	io.WriteString(w, sh.message)
}

func main() {
	err := http.ListenAndServe(":5000", StringHandler{
		"Hello World!",
	})
	if err != nil {
		Printfln("Error: %v", err.Error())
	}
}

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}
