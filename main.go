package main

import (
	"fmt"
	"net/http"

	"github.com/fjukstad/luft/controllers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", controllers.IndexHandler)
	mux.HandleFunc("/aqis/", controllers.AquisGeoJSON)
	mux.HandleFunc("/logs/", controllers.LogHandler)

	http.Handle("/public", http.FileServer(http.Dir("public")))

	fmt.Println("Server started on localhost:8000")
	err := http.ListenAndServe(":8000", mux)

	if err != nil {
		fmt.Println(err)
		return
	}

}
