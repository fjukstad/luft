package controllers

import (
	"encoding/csv"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/fjukstad/luftkvalitet"
	"github.com/paulmach/go.geojson"
)

const timeLayout = "2006-01-02T15:04:05.000Z"

func AquisGeoJSON(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	to, from, err := parseTimeInput(values)
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

func parseTimeInput(values url.Values) (to time.Time, from time.Time, err error) {
	to, err = time.Parse(timeLayout, values["to"][0])
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	from, err = time.Parse(timeLayout, values["from"][0])
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	return to, from, nil
}

// Return historical data for
// Results: station name, from time, to time, value, unit,
func HistoricalHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	to, from, err := parseTimeInput(values)
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
	records := [][]string{}
	header := []string{"station", "from", "to", "value", "component", "unit"}
	records = append(records, header)

	for _, hist := range historical {
		for _, m := range hist.Measurements {
			from := m.FromTime.Format(timeLayout)
			to := m.ToTime.Format(timeLayout)
			value := strconv.FormatFloat(m.Value, 'f', -1, 64)
			record := []string{hist.Station.Station, from, to, value, hist.Component, m.Unit}
			records = append(records, record)
		}
	}
	writer := csv.NewWriter(w)
	err = writer.WriteAll(records)
	if err != nil {
		w.Write([]byte("Could not write csv"))
		return
	}
}

// Get the first (current) air quality forecast for a specific area
func ForecastHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()

	f := luftkvalitet.Filter{
		Areas: values["area"],
	}

	forecasts, err := luftkvalitet.GetForecasts(f)
	if err != nil || len(forecasts) == 0 {
		w.Write([]byte("Fant ikke luftkvalitetsvarsel fra luftkvalitet.info"))
		return
	}

	w.Write([]byte(forecasts[0].Today[0].Description))

}
