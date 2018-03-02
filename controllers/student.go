package controllers

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
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
	
	var within string = ""
	var area string = ""

	if len(values["within"]) > 0 {
		within = values["within"][0]
	}
	if len(values["area"]) > 0 {
		area = values["area"][0]
	}

	filter := StudentFilter{
			ToTime:     to,
			FromTime:   from,
			Within:			within,
			Area:			  area,
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
		f.SetProperty("pmTen", measurement.PmTen)
		f.SetProperty("pmTwoFive", measurement.PmTwoFive)
		f.SetProperty("humidity", measurement.Humidity)
		f.SetProperty("temperature", measurement.Temperature)
		f.SetProperty("color", "6ee86e")
		f.SetProperty("weight", 10)
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

type StudentFilter struct {
	Area 				string
	Within 			string
	FromTime 		time.Time
	ToTime 			time.Time
}

func StudentHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	to, from, err := parseTimeInput(values)

	if err != nil {
		http.Error(w, "Could not parse time: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var within string = ""
	var area string = ""

	if len(values["within"]) > 0 {
		within = values["within"][0]
	}
	if len(values["area"]) > 0 {
		area = values["area"][0]
	}

	filter := StudentFilter{
			ToTime:     to,
			FromTime:   from,
			Within:			within,
			Area:			  area,
	}	

	data, err := getStudentData(filter)

	if err != nil {
		http.Error(w, "Could not parse student data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	records := [][]string{}
	header := []string{"timestamp", "latitude", "longitude", "pmTen", "pmTwoFive", "humidity", "temperature"}
	records = append(records, header)

	for _, measurement := range data {
		var latitude 				 float64
		var longitude 			 float64
		var valuePmTen 			 float64
		var valuePmTwoFive 	 float64
		var valueHumidity		 float64
		var valueTemperature float64
		
		latitude 					= measurement.Latitude
		longitude 				= measurement.Longitude
		valuePmTen 				= measurement.PmTen
		valuePmTwoFive 		= measurement.PmTwoFive
		valueHumidity 		= measurement.Humidity
		valueTemperature 	= measurement.Temperature

		formattedLatitude 				:= strconv.FormatFloat(latitude, 'f', -1, 64)
		formattedLongitude 				:= strconv.FormatFloat(longitude, 'f', -1, 64)
		formattedPmTenValue 			:= strconv.FormatFloat(valuePmTen, 'f', -1, 64)
		formattedPmTwoFiveValue 	:= strconv.FormatFloat(valuePmTwoFive, 'f', -1, 64)
		formattedHumidityValue		:= strconv.FormatFloat(valueHumidity, 'f', -1, 64)
		formattedTemperatureValue := strconv.FormatFloat(valueTemperature, 'f', -1, 64)

		timestamp := measurement.Date.Format(studentResponseTimeLayout)
		// station := measurement.Group

		record := []string{
							timestamp, 
							formattedLatitude,
							formattedLongitude,
							formattedPmTenValue,  
							formattedPmTwoFiveValue, 
							formattedHumidityValue,  
							formattedTemperatureValue, 
						}
		records = append(records, record)
	}

	writer := csv.NewWriter(w)

	filename := "studentdata.csv"
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)

	err = writer.WriteAll(records)
	if err != nil {
		http.Error(w, "Could not write csv", http.StatusInternalServerError)
		return
	}
}

var studentTimeLayout string = "2006-01-02T15:04:05"
var studentResponseTimeLayout string = "2006-01-02 15:04:05 -0700"

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
func getStudentData(filter StudentFilter) ([]Measurement, error) {

	fromDate := filter.FromTime.Format(studentTimeLayout)
	toDate := filter.ToTime.Format(studentTimeLayout)
	within := filter.Within
	area := filter.Area

	var u string
	if len(within) > 0 {
		u = "http://localhost:8080/api/data?totime=" + toDate + "&fromtime=" + fromDate + "&within=" + within
	}	else if len(area) > 0 {
		u = "http://localhost:8080/api/data?totime=" + toDate + "&fromtime=" + fromDate + "&area=" + url.QueryEscape(area)
	}	else {
		u = "http://localhost:8080/api/data?totime=" + toDate + "&fromtime=" + fromDate
	}
	// if len(within) > 0 {
	// 	u = "https://luft-184208.appspot.com/api/data?totime=" + toDate + "&fromtime=" + fromDate + "&within=" + within
	// }	else if len(area) > 0 {
	// 	u = "https://luft-184208.appspot.com/api/data?totime=" + toDate + "&fromtime=" + fromDate + "&area=" + url.QueryEscape(area)
	// }	else {
	// 	u = "https://luft-184208.appspot.com/api/data?totime=" + toDate + "&fromtime=" + fromDate
	// }

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

	for _, record := range records {
		if len(record) < 6 {
			return []Measurement{}, errors.Wrap(err, "error prasing csv, not enough records")
		}


		long, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			return []Measurement{}, errors.Wrap(err, "error parsing float (latitude)")
		}
		lat, err := strconv.ParseFloat(record[1], 64)
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
