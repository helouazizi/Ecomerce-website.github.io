// Ecomerce-website.github.io/back-end/hundler/hundler.go
package handler

import (
	"html/template"
	"log"
	"net/http"
)

// Define your page templates
type Pages struct {
	Home    *template.Template
	Contact *template.Template
	Account *template.Template
	SignIn  *template.Template
	Error   *template.Template
	About   *template.Template
	Cart    *template.Template
}

var page Pages

// Initialize templates
func init() {
	page.Home = parseTemplate("../front-end/pages/home.html")
	page.Error = parseTemplate("../front-end/pages/errors.html")
	page.About = parseTemplate("../front-end/pages/about.html")
	page.Account = parseTemplate("../front-end/pages/acount.html")
	page.SignIn = parseTemplate("../front-end/pages/sign_in.html")
	page.Cart = parseTemplate("../front-end/pages/cart.html")
	page.Contact = parseTemplate("../front-end/pages/contact.html")
}

// Function to parse a template
func parseTemplate(filename string) *template.Template {
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		log.Fatalf("Error parsing template file %s: %v", filename, err)
	}
	return tmpl
}

// Render a template
func renderTemplate(w http.ResponseWriter, tmpl *template.Template, data interface{}) {
	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Handlers for different pages
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		renderTemplate(w, page.Error, "404 NOT FOUND")
		return
	}
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		renderTemplate(w, page.Error, "405 METHOD NOT ALLOWED")
		return
	}
	renderTemplate(w, page.Home, nil)
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, page.Contact, nil)
}

func AccountHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, page.Account, nil)
}

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, page.SignIn, nil)
}

func CartHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, page.Cart, nil)
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, page.About, nil)
}

func CssHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		renderTemplate(w, page.Error, "405 METHOD NOT ALLOWED")
		return
	}

	http.ServeFile(w, r, "../front-end/css/style.css")
}
