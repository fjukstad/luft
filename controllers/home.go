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

var liveTemplate = template.Must(template.ParseFiles("views/base.html",
	"views/header.html", "views/navbar.html",
	"views/live.html", "views/footer.html"))

func LiveHandler(w http.ResponseWriter, r *http.Request) {
	err := liveTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}

var historyTemplate = template.Must(template.ParseFiles("views/base.html",
	"views/header.html", "views/navbar.html",
	"views/history.html", "views/footer.html"))

func HistoryHandler(w http.ResponseWriter, r *http.Request) {
	err := historyTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}
