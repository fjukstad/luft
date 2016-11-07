package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"

	"github.com/fjukstad/luftkvalitet"
	"github.com/gorilla/mux"
	"github.com/paulmach/go.geojson"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/aqis/{id}", AquisGeoJSON)
	r.HandleFunc("/public/{folder}/{file}", PublicHandler)

	fmt.Println("Server started on localhost:8000")
	err := http.ListenAndServe(":8000", r)

	if err != nil {
		fmt.Println(err)
		return
	}

}

//    "type": "Feature",
//    "properties": {
//        "name": "Coors Field",
//        "amenity": "Baseball Stadium",
//        "popupContent": "This is where the Rockies play!"
//    },
//    "geometry": {
//        "type": "Point",
//        "coordinates": [-104.99404, 39.75621]
//    }

func AquisGeoJSON(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
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
