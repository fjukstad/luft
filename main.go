package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/fjukstad/luft/controllers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/aqis", controllers.AquisGeoJSON)
	mux.HandleFunc("/historical", controllers.HistoricalHandler)
	mux.HandleFunc("/forecast", controllers.ForecastHandler)
	mux.HandleFunc("/logs", controllers.LogHandler)
	mux.HandleFunc("/student", controllers.StudentHandler)

	mux.Handle("/public/", http.FileServer(http.Dir(".")))

	mux.HandleFunc("/", controllers.IndexHandler)
	mux.HandleFunc("/om", controllers.AboutHandler)
	mux.HandleFunc("/historikk", controllers.HistoryHandler)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	fmt.Println("Server started on port", port)
	err := http.ListenAndServe(":"+port, mux)

	if err != nil {
		fmt.Println(err)
		return
	}

}
