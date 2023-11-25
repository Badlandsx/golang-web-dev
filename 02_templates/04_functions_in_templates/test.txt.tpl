Here is the team
// Functions template.FuncMap or predefined functions
{{range .Players}}- name: {{upper .Name}} - poste: {{.Position}} {{ if eq .Name "Luka Doncic" }}(The best player in the world){{end}}{{printf "\n"}}{{end}}

// Representation of the team (via functions)
{{ . | showTeam }}

// Via method on team struct
{{ .Show }}

// format date with function
{{ formatDate .Time }}

// Predefined functions lt (lower than)
{{ if lt 4 8}}Test1{{end}}

// Define variable and use it
{{ define "tpl1"}}Denver champion NBA 2023{{ end }}
{{ template "tpl1" }}