package main

import (
	"html/template"
	"os"
)

type PersonA struct {
	Name 	string
	Sex 	string
}


func main(){
	P := PersonA{
		Name: "云无月",
		Sex: "女",
	}
	tmpl3, _ := template.New("test").Parse("Name: {{.Name}}, Age: {{.Age}}")
	_ = tmpl3.Execute(os.Stdout, P)
}