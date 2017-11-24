package controllers

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/fjukstad/luftkvalitet"
	"github.com/paulmach/go.geojson"
	"github.com/pkg/errors"
)

func StudentAqisHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	to, from, err := parseTimeInput(values)
	if err != nil {
		http.Error(w, "Could not parse time: "+err.Error(), http.StatusInternalServerError)
		return
	}

	component := values["component"][0]

	filter := luftkvalitet.Filter{
		ToTime:     to,
		FromTime:   from,
		Components: []string{component},
	}

	data, err := getStudentData(filter)
	if err != nil {
		http.Error(w, "Could not parse student data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fc := geojson.NewFeatureCollection()

	for _, measurement := range data {
		geom := geojson.NewPointGeometry([]float64{measurement.Longitude, measurement.Latitude})
		f := geojson.NewFeature(geom)
		f.SetProperty("date", measurement.Date)
		f.SetProperty("dust", measurement.PmTen)
		f.SetProperty("humidity", measurement.Humidity)
		f.SetProperty("temperature", measurement.Temperature)
		f.SetProperty("weight", 2)
		fc = fc.AddFeature(f)
	}

	b, err := fc.MarshalJSON()
	if err != nil {
		http.Error(w, "Could not marshal geojson"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(b)
	return
}

func StudentHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	to, from, err := parseTimeInput(values)
	if err != nil {
		http.Error(w, "Could not parse time: "+err.Error(), http.StatusInternalServerError)
		return
	}

	component := values["component"][0]


	filter := luftkvalitet.Filter{
		ToTime:     to,
		FromTime:   from,
		Components: []string{component},
	}
	data, err := getStudentData(filter)
	if err != nil {
		http.Error(w, "Could not parse student data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	records := [][]string{}
	header := []string{"from", "to", "value", "component", "unit"}
	records = append(records, header)

	for _, measurement := range data {
		var value float64
		var unit string

		switch component {
		case "dust":
			value = measurement.PmTen
			unit = "ug/m3"
		case "humidity":
			value = measurement.Humidity
			unit = "%"
		case "temperature":
			value = measurement.Temperature
			unit = "C"
		}

		formattedValue := strconv.FormatFloat(value, 'f', -1, 64)

		from := measurement.Date.Format(timeLayout)
		to := measurement.Date.Format(timeLayout)

		// station := measurement.Group

		record := []string{from, to, formattedValue, component, unit}
		records = append(records, record)
	}

	writer := csv.NewWriter(w)

	filename := "student-" + component + ".csv"
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)

	err = writer.WriteAll(records)
	if err != nil {
		http.Error(w, "Could not write csv", http.StatusInternalServerError)
		return
	}
}

var studentTimeLayout string = "2006-01-02T15:04:05"
var studentResponseTimeLayout string = "2006-01-02 15:04:05 +0100"

type Measurement struct {
	Latitude      float64
	Longitude     float64
	PmTen         float64
	PmTwoFive     float64
	Humidity      float64
	Temperature   float64
	Date          time.Time
}

// Fetches and parses the student collected data
func getStudentData(filter luftkvalitet.Filter) ([]Measurement, error) {

	fromDate := filter.FromTime.Format(studentTimeLayout)
	toDate := filter.ToTime.Format(studentTimeLayout)

	u := "http://localhost:8080/api/data?totime=" + toDate + "&fromtime=" + fromDate
	// u := "https://luft-184208.appspot.com/download?totime=" + toDate + "&fromtime=" + fromDate

	resp, err := http.Get(u)
	if err != nil {
		return []Measurement{}, errors.Wrap(err, "Could not download data from luftprosjekttromso")
	}

	reader := csv.NewReader(resp.Body)

	records, err := reader.ReadAll()
	if err != nil {
		if len(records) == 0 {
			return []Measurement{}, nil
		}
		return []Measurement{}, errors.Wrap(err, "Could not read csv from "+ u)
	}

	//fc := geojson.NewFeatureCollection()
	var data []Measurement

	for i, record := range records {
		// skipping header
		if i == 0 {
			continue
		}

		if len(record) < 6 {
			return []Measurement{}, errors.Wrap(err, "error prasing csv, not enough records")
		}


		lat, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			return []Measurement{}, errors.Wrap(err, "error parsing float (latitude)")
		}
		long, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			return []Measurement{}, errors.Wrap(err, "error parsing float (longitude)")
		}

		humid, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			return []Measurement{}, errors.Wrap(err, "error parsing float (humidity)")
		}

		temp, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			return []Measurement{}, errors.Wrap(err, "error parsing float (temperature)")
		}
		pmTen, err := strconv.ParseFloat(record[4], 64)
		if err != nil {
			return []Measurement{}, errors.Wrap(err, "error parsing float (pmTen)")
		}
		pmTwoFive, err := strconv.ParseFloat(record[5], 64)
		if err != nil {
			return []Measurement{}, errors.Wrap(err, "error parsing float (pmTwoFive)")
		}

		date, err := time.Parse(studentResponseTimeLayout, record[6])
		if err != nil {
			fmt.Println(err)
			msg := "Could not parse date " + record[6] + " skipping measurement.\n"
			msg += "Url: " + u
			fmt.Println(msg)
			fmt.Println("Full record: ", record)
			continue
			//return []Measurement{}, errors.Wrap(err, msg)
		}


		data = append(data, Measurement{
			lat,
			long,
			pmTen,
			pmTwoFive,
			humid,
			temp,
			date,
		})
	}

	return data, nil

}
