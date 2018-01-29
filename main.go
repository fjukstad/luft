package main

import (
	"fmt"
	"net/http"
	"os"
	// "github.com/fjukstad/luft/controllers"
	"github.com/luft-1/controllers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/niluaqis", controllers.NILUAqiHandler)
	mux.HandleFunc("/historical", controllers.HistoricalHandler)
	mux.HandleFunc("/forecast", controllers.ForecastHandler)
	mux.HandleFunc("/studentaqis", controllers.StudentAqisHandler)
	mux.HandleFunc("/student", controllers.StudentHandler)
	mux.HandleFunc("/precipitation", controllers.PrecipitationHandler)
	mux.HandleFunc("/sendfile", controllers.PostFileHandler)

	mux.Handle("/public/", http.FileServer(http.Dir(".")))

	mux.HandleFunc("/", controllers.IndexHandler)
	mux.HandleFunc("/live", controllers.LiveHandler)
	mux.HandleFunc("/historikk", controllers.HistoryHandler)
	mux.HandleFunc("/lastopp", controllers.UploadHandler)

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
