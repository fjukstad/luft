package controllers

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/paulmach/go.geojson"
)

func LogHandler(w http.ResponseWriter, r *http.Request) {
	filename := "data/DATA2.csv"
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
	for i, record := range records {
		// skipping header
		if i == 0 {
			continue
		}
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
