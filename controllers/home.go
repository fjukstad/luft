package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
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

func PublicHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	folder := vars["folder"]
	file := vars["file"]

	base := "public/"
	filename := base + folder + "/" + file

	f, err := ioutil.ReadFile(filename)
	if err != nil {
		w.Write([]byte("Could not find file " + filename))
	} else {
		w.Write(f)
	}
}
