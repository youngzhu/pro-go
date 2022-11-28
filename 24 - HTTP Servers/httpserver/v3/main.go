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
	Printfln("Request for %v", r.URL.Path)
	switch r.URL.Path {
	case "/favicon.ico":
		http.NotFound(w, r)
	case "/message":
		io.WriteString(w, sh.message)
	default:
		http.Redirect(w, r, "/message", http.StatusTemporaryRedirect)
	}
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
