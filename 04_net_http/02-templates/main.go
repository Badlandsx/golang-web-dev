package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.html.tpl"))
}

func TemplateHandler(w http.ResponseWriter, req *http.Request) {
	// Allow to get params (from GET url params or POST form)
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	// Possibility to set header
	w.Header().Set("Test", "Test-header")
	tpl.ExecuteTemplate(w, "index.html.tpl", req)
}

func main() {
	http.HandleFunc("/", TemplateHandler)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
