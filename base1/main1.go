package main

import (
"html/template"
"net/http"
)

type Person struct {
	Name 	string
	Sex 	string
}

func tmpl2(w http.ResponseWriter, r *http.Request){
	t1, err := template.ParseFiles("./test.html")
	if err != nil {
		panic(err)
	}
	t1.Execute(w, Person{
		Name: "zhou",
		Sex: "男",
	})
}


func main(){
	server := http.Server{
		Addr: "localhost:9999",
	}
	http.HandleFunc("/tmpl", tmpl2)
	server.ListenAndServe()
}