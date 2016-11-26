package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"text/template"

	"github.com/fjukstad/luftkvalitet"
	"github.com/gorilla/mux"
	"github.com/paulmach/go.geojson"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", IndexHandler)
	mux.HandleFunc("/aqis/", AquisGeoJSON)
	mux.HandleFunc("/logs/", LogHandler)

	http.Handle("/public", http.FileServer(http.Dir("public")))

	fmt.Println("Server started on localhost:8000")
	err := http.ListenAndServe(":8000", mux)

	if err != nil {
		fmt.Println(err)
		return
	}

}

func LogHandler(w http.ResponseWriter, r *http.Request) {
	filename := "data/test-data.csv"
	f, err := os.Open(filename)
	if err != nil {
		w.Write([]byte("Could not open data file " + err.Error()))
		return
	}

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		w.Write([]byte("Could not parse csv file " + err.Error()))
		return
	}

	fc := geojson.NewFeatureCollection()
	for _, record := range records {
		if len(record) < 6 {
			w.Write([]byte("error parsing csv"))
			return
		}

		date := record[0]
		lat, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			w.Write([]byte("error parsing float " + err.Error()))
			return
		}
		long, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			w.Write([]byte("error parsing float " + err.Error()))
			return
		}
		ppm := record[3]
		humid := record[4]
		temp := record[5]

		geom := geojson.NewPointGeometry([]float64{long, lat})
		f := geojson.NewFeature(geom)
		f.SetProperty("date", date)
		f.SetProperty("ppm", ppm)
		f.SetProperty("humid", humid)
		f.SetProperty("temp", temp)
		fc = fc.AddFeature(f)
	}
	b, err := fc.MarshalJSON()

	if err != nil {
		fmt.Println(err)
		w.Write([]byte("Could not marshal geojson"))
		return
	}

	w.Write(b)
}

func AquisGeoJSON(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Path
	id := strings.Trim(p, "/aqis/")
	id = "{" + id + "}"

	aqis, err := luftkvalitet.GetAqis(id)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("could not get data from luftkvalutet"))
		return
	}

	fc := geojson.NewFeatureCollection()
	for _, aqi := range aqis {
		geom := geojson.NewPointGeometry([]float64{aqi.Lon, aqi.Lat})
		f := geojson.NewFeature(geom)
		f.SetProperty("name", aqi.Name)
		f.SetProperty("popupContent", aqi.Description)
		f.SetProperty("color", aqi.Color)
		fc = fc.AddFeature(f)
	}

	b, err := fc.MarshalJSON()

	if err != nil {
		fmt.Println(err)
		w.Write([]byte("Could not marshal geojson"))
		return
	}

	w.Write(b)
	return
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
