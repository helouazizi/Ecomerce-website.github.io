// Ecomerce-website.github.io/back-end/main.go
package main

import (
	handler "Ecomerce/hundler"
	"fmt"
	"log"
	"net/http"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//log.Printf("Received request: %s %s", r.Method, r.URL.Path) // Log the request method and path
		next.ServeHTTP(w, r) // Call the next handler in the chain
	})
}

func main() {

	http.Handle("/imges/", http.StripPrefix("/imges/", http.FileServer(http.Dir("../front-end/imges"))))
	http.Handle("/", loggingMiddleware(http.HandlerFunc(handler.HomeHandler)))
	http.Handle("/contact", loggingMiddleware(http.HandlerFunc(handler.ContactHandler)))
	http.Handle("/account", loggingMiddleware(http.HandlerFunc(handler.AccountHandler)))
	http.Handle("/sign_in", loggingMiddleware(http.HandlerFunc(handler.SignInHandler)))
	http.Handle("/cart", loggingMiddleware(http.HandlerFunc(handler.CartHandler)))
	http.Handle("/about", loggingMiddleware(http.HandlerFunc(handler.AboutHandler)))
	http.Handle("/css/style.css", loggingMiddleware(http.HandlerFunc(handler.CssHandler)))
	fmt.Println("server listnning  on port 8080 >>> http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
