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
	as := strings.Trim(p, "/aqis/")
	areas := strings.Split(as, ",")

	f := luftkvalitet.Filter{
		Areas: areas,
	}

	measurements, err := luftkvalitet.GetMeasurements(f)
	if err != nil {
		w.Write([]byte("fuck"))
		return
	}

	fc := geojson.NewFeatureCollection()
	for _, m := range measurements {
		geom := geojson.NewPointGeometry([]float64{m.Longitude, m.Latitude})
		f := geojson.NewFeature(geom)
		f.SetProperty("name", m.Station.Station)
		f.SetProperty("component", m.Component)
		f.SetProperty("unit", m.Unit)
		f.SetProperty("value", m.Value)
		//f.SetProperty("popupContent", aqi.Description)
		f.SetProperty("color", m.Color)
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
