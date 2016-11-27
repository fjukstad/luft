package luftkvalitet

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
)

type Measurement struct {
	Location     string
	Component    string
	Time         string
	Value        string
	Date         string
	DailyAverage string
	Unit         string
}

type Aqi struct {
	Name             string
	Index            int
	Color            string
	Text             string
	ShortDescription string
	Description      string
	Lat              float64
	Lon              float64
	Url              string
	Type             int
}

type AqiRequest struct {
	Culture string `json:"culture"`
	Type    string `json:"type"`
	Id      string `json:"id"`
}

type AqiResponse struct {
	D AqiData `json:"d"`
}

type AqiData struct {
	UnkownType string `json:"__type"`
	Type       string
	Aqis       []Aqi
}

func GetAqis(id string) ([]Aqi, error) {
	url := "http://luftkvalitet.info/Nilu/Webservice/CityService.asmx/GetAqis"

	ar := AqiRequest{"no", "1", id}
	b, err := json.Marshal(ar)
	if err != nil {
		return []Aqi{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []Aqi{}, err
	}

	body, _ := ioutil.ReadAll(resp.Body)
	aqiResponse := AqiResponse{}

	err = json.Unmarshal(body, &aqiResponse)
	if err != nil {
		return []Aqi{}, err
	}

	return aqiResponse.D.Aqis, nil
}

func GetMeasurements(location string) ([]Measurement, error) {

	resp, err := http.Get("http://luftkvalitet.info/home/overview.aspx?type=Station&id={" + location + "}")

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return []Measurement{}, errors.Wrap(err, "Could not parse luftkvalitet.info")
	}

	measurements := []Measurement{}

	doc.Find("table#ctl00_cph_Map_ctl00_gwStation").Each(func(i int, s *goquery.Selection) {
		s.Find("tr").Each(func(i int, s *goquery.Selection) {
			if s.HasClass("tableHead") {
				return
			} else {
				measurement := Measurement{}
				measurement.Location = doc.Find("#ctl00_cph_Text_ctl00_lTitle").Text()
				s.Find("td").Each(func(i int, s *goquery.Selection) {
					text := s.Text()
					text = strings.TrimSpace(text)
					switch i {
					case 0:
						measurement.Component = text
					case 1:
						measurement.Time = text
					case 2:
						measurement.Value = text
					case 3:
						measurement.Date = text
					case 4:
						measurement.DailyAverage = text
					case 5:
						measurement.Unit = text
					}
				})
				measurements = append(measurements, measurement)
			}
		})

	})

	return measurements, nil

}
