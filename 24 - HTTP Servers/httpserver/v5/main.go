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
	io.WriteString(w, sh.message)
}

func main() {
	http.Handle("/message", StringHandler{"Hello World!"})
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/", http.RedirectHandler("/message", http.StatusTemporaryRedirect))

	go func() {
		err := http.ListenAndServeTLS(":5500", "certificate.cer",
			"certificate.key", nil)
		if err != nil {
			Printfln("HTTPS Error: %v", err.Error())
		}
	}()

	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		Printfln("Error: %v", err.Error())
	}
}

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}
