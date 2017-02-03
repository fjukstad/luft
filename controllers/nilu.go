package controllers

import (
	"net/http"
	"time"

	"github.com/fjukstad/luftkvalitet"
	"github.com/paulmach/go.geojson"
)

const timeLayout = "2006-01-02T15:04:05.000Z"

func AquisGeoJSON(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()

	to, err := time.Parse(timeLayout, values["to"][0])
	if err != nil {
		w.Write([]byte("Could not parse time" + err.Error()))
		return
	}

	from, err := time.Parse(timeLayout, values["from"][0])
	if err != nil {
		w.Write([]byte("Could not parse time" + err.Error()))
		return
	}

	components := values["component"]

	areas := values["area"]

	f := luftkvalitet.Filter{
		Areas:      areas,
		ToTime:     to,
		FromTime:   from,
		Components: components,
	}

	historical, err := luftkvalitet.GetHistorical(f)
	if err != nil {
		w.Write([]byte("could not get data from api.nilu.no."))
		return
	}

	fc := geojson.NewFeatureCollection()
	for _, hist := range historical {
		geom := geojson.NewPointGeometry([]float64{hist.Location.Longitude, hist.Location.Latitude})
		for _, m := range hist.Measurements {
			f := geojson.NewFeature(geom)
			f.SetProperty("name", hist.Station.Station)
			f.SetProperty("component", hist.Component)
			f.SetProperty("unit", m.Unit)
			f.SetProperty("value", m.Value)
			f.SetProperty("color", m.Color)
			fc = fc.AddFeature(f)
		}
	}

	b, err := fc.MarshalJSON()
	if err != nil {
		w.Write([]byte("Could not marshal geojson " + err.Error()))
		return
	}

	w.Write(b)
	return
}
