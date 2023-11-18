package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	var tpl *template.Template
	var err error
	tpl, err = template.ParseFiles("model.go.tpl", "simple.txt.tpl")
	if err != nil {
		log.Fatalln(err)
	}

	// Execute the first one template
	if err := tpl.Execute(os.Stdout, nil); err != nil {
		log.Fatalln(err)
	}

	// Execute specific simple.txt.tpl template
	log.Printf("\n\nExecute template simple.txt.tpl\n")
	if err := tpl.ExecuteTemplate(os.Stdout, "simple.txt.tpl", nil); err != nil {
		log.Fatalln(err)
	}

	// ParseGlob and Must (error checking) -> this should be used with function init()
	/*
			    var tpl *template.Template
		        func init() {
			        tpl = template.Must(template.ParseGlob("templates/*"))
		        }
	*/
	tpl2 := template.Must(template.ParseGlob("*.tpl"))
	log.Printf("Template 2 parseGlob tree: %v", tpl2.Tree)
}
