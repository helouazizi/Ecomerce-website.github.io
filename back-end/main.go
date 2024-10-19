// Ecomerce-website.github.io/back-end/main.go
package main

import (
	handler "Ecomerce/hundler"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.HomeHandler)
	http.HandleFunc("/contact", handler.ContactHandler)
	http.HandleFunc("/account", handler.AccountHandler)
	http.HandleFunc("/signin", handler.SignInHandler)
	http.HandleFunc("/about", handler.AboutHandler)
	http.HandleFunc("/css/", handler.CssHandler)
	fmt.Println("server listnning  on port 8080 >>> http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
