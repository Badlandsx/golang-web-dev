package model

type {{ .TypeName }} struct {
    Name string
}

func GetName(t {{ .TypeName }}) string {
    return t.Name
}