package main

import (
	"html/template"
	"net/http"
)

func tmpl(w http.ResponseWriter, r *http.Request){
	t1, err := template.ParseFiles("./test.html")
	if err != nil {
		panic(err)
	}
	t1.Execute(w, "hello world")
}

func main(){
	server := http.Server{
		Addr: "localhost:9999",
	}
	http.HandleFunc("/tmpl", tmpl)
	server.ListenAndServe()
}