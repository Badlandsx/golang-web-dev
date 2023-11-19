package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("model1/*.tpl"))
}

func main() {
	//1
	simpleDataTpl()

	//2 with struc
	teamDataTpl()
}

// model1
func simpleDataTpl() {

	packageName := "model1"

	// Creation of package directory
	if _, err := os.Stat(packageName); os.IsNotExist(err) {
		err := os.Mkdir(packageName, 0660)
		if err != nil {
			log.Fatalln(err)
		}
	}

	// Creation of file
	file, err := os.Create(fmt.Sprintf("%s/%s.go", packageName, packageName))
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(file, "User")
	if err != nil {
		log.Fatalln(err)
	}
}

// model2
type Player struct {
	Name     string
	Position int
}

type Team struct {
	Name    string
	Players []Player
}

func teamDataTpl() {

	packageName := "model2"

	// Creation of package directory
	if _, err := os.Stat(packageName); os.IsNotExist(err) {
		err := os.Mkdir(packageName, 0660)
		if err != nil {
			log.Fatalln(err)
		}
	}

	// Creation of file
	file, err := os.Create(fmt.Sprintf("%s/%s.txt", packageName, packageName))
	if err != nil {
		log.Fatalln(err)
	}

	tpl2, err := template.ParseGlob(packageName + "/*.tpl")
	if err != nil {
		log.Fatalln(err)
	}

	dallas := Team{
		"Dallas Mavericks",
		[]Player{{
			"Luka Doncic",
			1,
		},
			{
				"Kyrie Irving",
				1,
			},
		},
	}

	err = tpl2.Execute(file, dallas)
	if err != nil {
		log.Fatalln(err)
	}
}
