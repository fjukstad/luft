package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/public/{folder}/{file}", PublicHandler)

	fmt.Println("Server started on localhost:8000")
	err := http.ListenAndServe(":8000", r)

	if err != nil {
		fmt.Println(err)
		return
	}

}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	indexTemplate := template.Must(template.ParseFiles("index.html"))
	indexTemplate.Execute(w, nil)

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
