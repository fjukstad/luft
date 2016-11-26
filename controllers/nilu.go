package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/fjukstad/luftkvalitet"
	"github.com/paulmach/go.geojson"
)

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
