// Ecomerce-website.github.io/back-end/main.go
package main

import (
	"Ecomerce/back-end/hundler"
	"fmt"
	"net/http"
	"os"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
	}
	fmt.Println("Current working directory:", dir)
	http.HandleFunc("/", hundler.HomeHundler)

	fmt.Println("server listnning  on port 8080 >>> http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
