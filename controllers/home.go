package controllers

import (
	"fmt"
	"net/http"
	"text/template"
)

var indexTemplate = template.Must(template.ParseFiles("views/base.html",
	"views/header.html", "views/navbar.html", "views/navbar2.html",
	"views/index.html", "views/footer.html", "views/scripts.html"))

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	err := indexTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}

var liveTemplate = template.Must(template.ParseFiles("views/base.html",
	"views/header.html", "views/navbar.html",
	"views/live.html", "views/footer.html", "views/scripts.html"))

func LiveHandler(w http.ResponseWriter, r *http.Request) {
	err := liveTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}

var historyTemplate = template.Must(template.ParseFiles("views/base.html",
	"views/header.html", "views/navbar.html",
	"views/history.html", "views/footer.html", "views/scripts.html"))

func HistoryHandler(w http.ResponseWriter, r *http.Request) {
	err := historyTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}

var uploadTemplate = template.Must(template.ParseFiles("views/base.html",
	"views/header.html", "views/navbar.html",
	"views/upload.html", "views/footer.html", "views/scripts.html"))

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	err := uploadTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}

var resourcesTemplate = template.Must(template.ParseFiles("views/base.html",
	"views/header.html", "views/navbar.html", "views/navbar2.html",
	"views/resources.html", "views/footer.html", "views/scripts.html"))

func ResourcesHandler(w http.ResponseWriter, r *http.Request) {
	err := resourcesTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}
