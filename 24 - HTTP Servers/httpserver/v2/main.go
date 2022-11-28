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
	if r.URL.Path == "/favicon.ico" {
		Printfln("Request for icon detected - return 404")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	Printfln("Request for %v", r.URL.Path)
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
