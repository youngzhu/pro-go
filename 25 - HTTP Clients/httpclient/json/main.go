package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func main() {
	go http.ListenAndServe(":5000", nil)
	time.Sleep(time.Second)

	resp, err := http.Get("http://localhost:5000/json")
	if err != nil {
		Printfln("Error: %v", err.Error())
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var data []Product
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err == nil {
			for _, p := range data {
				Printfln("Name: %v, Price: $%.2f", p.Name, p.Price)
			}
		} else {
			Printfln("Decode error: %v", err.Error())
		}
	} else {
		Printfln("Status Code: %v", resp.StatusCode)
	}
}
