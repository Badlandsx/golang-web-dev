package model

type {{ .}} struct {
    Name string
    Description string
}

func GetName(t *{{.}}) string {
    return t.Name
}
{{$variable1 := .}}
func GetDescription(t *{{$variable1}}) string {
    return t.Description
}
