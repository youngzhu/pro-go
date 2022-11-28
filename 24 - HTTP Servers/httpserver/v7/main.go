package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type StringHandler struct {
	message string
}

func (sh StringHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	Printfln("Request for %v", r.URL.Path)
	io.WriteString(w, sh.message)
}

func HTTPSRedirect(w http.ResponseWriter, r *http.Request) {
	host := strings.Split(r.Host, ":")[0]
	target := "https://" + host + ":5500" + r.URL.Path
	if len(r.URL.RawQuery) > 0 {
		target += "?" + r.URL.RawQuery
	}
	http.Redirect(w, r, target, http.StatusTemporaryRedirect)
}

func main() {
	http.Handle("/message", StringHandler{"Hello World!"})
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/", http.RedirectHandler("/message", http.StatusTemporaryRedirect))

	fsHandler := http.FileServer(http.Dir("./static"))
	http.Handle("/files/", http.StripPrefix("/files", fsHandler))

	go func() {
		err := http.ListenAndServeTLS(":5500", "certificate.cer",
			"certificate.key", nil)
		if err != nil {
			Printfln("HTTPS Error: %v", err.Error())
		}
	}()

	err := http.ListenAndServe(":5000",
		http.HandlerFunc(HTTPSRedirect))
	if err != nil {
		Printfln("Error: %v", err.Error())
	}
}

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}
