package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
	"time"
)

type Player struct {
	Name     string
	Position int
}

type Team struct {
	Time    time.Time
	Name    string
	Players []Player
}

func (t Team) Show() string {
	return fmt.Sprintf(`
	  Name: %s,
	  Players: %v
	`, t.Name, t.Players)
}

func TimeFormat(t time.Time) string {
	return t.Format("2023-12-22")
}

func ShowTeam(t Team) string {
	return fmt.Sprintf(`
	  Name: %s,
	  Players: %v
	`, t.Name, t.Players)
}

var tpl *template.Template
var tplFuncMaps = template.FuncMap{
	"upper":      strings.ToUpper,
	"showTeam":   ShowTeam,
	"formatDate": TimeFormat,
}

func init() {
	tpl = template.Must(template.New("test").Funcs(tplFuncMaps).ParseFiles("test.txt.tpl"))
}

func main() {
	// Creation of file
	file, err := os.Create("test.txt")
	if err != nil {
		log.Fatalln(err)
	}

	dallas := Team{
		time.Now(),
		"Dallas Mavericks",
		[]Player{
			{
				"Luka Doncic",
				1,
			},
			{
				"Kyrie Irving",
				1,
			},
		},
	}

	err = tpl.ExecuteTemplate(file, "test.txt.tpl", dallas)
	if err != nil {
		log.Fatalln(err)
	}
}
