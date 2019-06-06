package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func templateHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/layout.html")
	fmt.Println(t.Name())
	t.Execute(w, "Hello world")
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/index", templateHandler)

}
