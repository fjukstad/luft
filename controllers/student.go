package controllers

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/fjukstad/luftkvalitet"
	"github.com/paulmach/go.geojson"
	"github.com/pkg/errors"
)

func LogHandler(w http.ResponseWriter, r *http.Request) {
	filename := "data/smaller.csv"
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
		f.SetProperty("group", "Student-A")
		f.SetProperty("date", date)
		f.SetProperty("value", ppm)
		f.SetProperty("component", "PM2.5")
		f.SetProperty("humid", humid)
		f.SetProperty("temp", temp)
		f.SetProperty("weight", 2)
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
		fmt.Println("Could not parse student data:" + err.Error())
		return
	}

	records := [][]string{}
	header := []string{"station", "from", "to", "value", "component", "unit"}
	records = append(records, header)

	for _, measurement := range data {
		var value float64
		var unit string

		switch component {
		case "dust":
			value = measurement.Dust
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

		station := measurement.Group

		record := []string{station, from, to, formattedValue, component, unit}
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

var studentTimeLayout string = "02.01.2006"
var studentResponseTimeLayout string = "2006-01-02 15:04:05 UTC"

type Measurement struct {
	Id            string
	Latitude      float64
	Longitude     float64
	Dust          float64
	Humidity      float64
	Temperature   float64
	SubmittedDate time.Time
	UpdatedDate   time.Time
	Group         string
	Date          time.Time
}

// Fetches and parses the student collected data
func getStudentData(filter luftkvalitet.Filter) ([]Measurement, error) {

	fromDate := filter.FromTime.Format(studentTimeLayout)
	toDate := filter.ToTime.Format(studentTimeLayout)

	u := "http://luftprosjekttromso.herokuapp.com/files/get_data?totime=" + toDate + "&fromtime=" + fromDate

	resp, err := http.Get(u)
	if err != nil {
		return []Measurement{}, errors.Wrap(err, "Could not download data from luftprosjekttromso")
	}

	reader := csv.NewReader(resp.Body)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println(resp.Body)
		return []Measurement{}, errors.Wrap(err, "Could not read csv from "+u)
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

		id := record[0]
		lat, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			return []Measurement{}, errors.Wrap(err, "error parsing float (latitude)")
		}
		long, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			return []Measurement{}, errors.Wrap(err, "error parsing float (longitude)")
		}
		dust, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			return []Measurement{}, errors.Wrap(err, "error parsing float (dust)")
		}

		humid, err := strconv.ParseFloat(record[4], 64)
		if err != nil {
			return []Measurement{}, errors.Wrap(err, "error parsing float (humidity)")
		}

		temp, err := strconv.ParseFloat(record[5], 64)
		if err != nil {
			return []Measurement{}, errors.Wrap(err, "error parsing float (temperature)")
		}

		submittedDate, err := time.Parse(studentResponseTimeLayout, record[6])
		if err != nil {
			fmt.Println("Could not parse submitted date.")
			fmt.Println(record[6])
			fmt.Println("Continuing.")
		}

		updatedDate, err := time.Parse(studentResponseTimeLayout, record[7])
		if err != nil {
			fmt.Println("Could not parse updated date.")
			fmt.Println(record[7])
			fmt.Println("Continuing.")
		}

		group := record[8]

		date, err := time.Parse(studentResponseTimeLayout, record[9])
		if err != nil {
			msg := "Could not parse date " + record[9]
			return []Measurement{}, errors.Wrap(err, msg)
		}

		data = append(data, Measurement{
			id,
			lat,
			long,
			dust,
			humid,
			temp,
			submittedDate,
			updatedDate,
			group,
			date,
		})
	}

	return data, nil

}
