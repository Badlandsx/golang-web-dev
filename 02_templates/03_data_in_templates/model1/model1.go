package model

type User struct {
    Name string
    Description string
}

func GetName(t *User) string {
    return t.Name
}

func GetDescription(t *User) string {
    return t.Description
}
