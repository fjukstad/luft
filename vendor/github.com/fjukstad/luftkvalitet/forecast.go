package luftkvalitet

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Result struct {
	Area
	Today    []Forecast `json:"today"`
	Tomorrow []Forecast `json:"tomorrow"`
}

type Forecast struct {
	Index        int    `json:"index"`
	Description  string `json:"description"`
	ForecastDate string `json:"forecastDate"`
	TimeOfDay    int    `json:"timeOfDay"`
}

func GetForecasts(f Filter) ([]Result, error) {
	url := endpoint + "aq/forecast.json"

	url = addFilter(url, f)

	resp, err := http.Get(url)

	if err != nil {
		return []Result{}, err
	}

	body, _ := ioutil.ReadAll(resp.Body)

	var result []Result
	err = json.Unmarshal(body, &result)
	if err != nil {
		return []Result{}, err
	}

	return result, nil

}
