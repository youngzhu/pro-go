package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleJsonRequest(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(Products)
}

func init() {
	http.HandleFunc("/json", HandleJsonRequest)
}

func main() {
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		Printfln("Error: %v", err.Error())
	}
}

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}
