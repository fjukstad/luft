package controllers

import (
	"fmt"
	"net/http"
	"text/template"
)

var indexTemplate = template.Must(template.ParseFiles("views/base.html",
	"views/header.html", "views/navbar.html",
	"views/index.html", "views/footer.html"))

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	err := indexTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}

var aboutTemplate = template.Must(template.ParseFiles("views/base.html",
	"views/header.html", "views/navbar.html",
	"views/about.html", "views/footer.html"))

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	err := aboutTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}
