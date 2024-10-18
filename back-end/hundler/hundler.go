// Ecomerce-website.github.io/back-end/hundler/hundler.go
package hundler

import (
	"html/template"
	"log"
	"net/http"
)

/*
create all  the routes for the application
and use them  to serve the templates
it depend  on the path of the url
*/
type Pages struct {
	ErrorTemplate   *template.Template
	Abouttemplate   *template.Template
	Sign_inTemplate *template.Template
	CartTemplate    *template.Template
	ContactTemplate *template.Template
	Acounttemplate  *template.Template
}

type Components struct {
	Header *template.Template
	Footer *template.Template
}

var page Pages
//var components Components

/*
 initialase  the templates to use them as needed
*/

func init() {
	page.ErrorTemplate = parseTemplate("../front-end/pages/home.html")
	/*page.Abouttemplate = parseTemplate("about.html")
	page.Acounttemplate = parseTemplate("acount.html")
	page.Sign_inTemplate = parseTemplate("sign_in.html")
	page.CartTemplate = parseTemplate("cart.html")
	page.ContactTemplate = parseTemplate("contact.html")
	components.Header = parseComponent("header.html")*/
	//components.Footer = parseComponent("footer.html")
}

func parseTemplate(filename string) *template.Template {
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		log.Fatalf("Error parsing template file %s: %v", filename, err)
	}
	return tmpl
}
/*
func parseComponent(filename string) *template.Template {
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		log.Fatal()
	}
	return tmpl
}*/

func HomeHundler(w http.ResponseWriter, r *http.Request) {
	/*if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		page.Abouttemplate.Execute(w, "")
		return
	}*/
	page.ErrorTemplate.Execute(w, nil)

}
