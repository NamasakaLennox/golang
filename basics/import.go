package main

import (
	"fmt"
	foo "net/http"
)

func main() {
	fmt.Println("Hello, Go Standard Lib")

	resp, err := foo.Get("https://www.google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("HTTP Response Status: ", resp.Status)
}
