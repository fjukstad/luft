package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/fjukstad/met"
)

var metTimeLayout = "2006-01-02T15:04:05.000Z"

func PrecipitationHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	stations := values["station"]
	elements := values["elements"]

	to, from, err := parseTimeInput(values)
	if err != nil {
		http.Error(w, "Could not parse time: "+err.Error(), http.StatusInternalServerError)
		return
	}

	f := met.Filter{
		Sources:       stations,
		ReferenceTime: to.Format(metTimeLayout) + "/" + from.Format(metTimeLayout),
		Elements:      elements,
	}

	data, err := met.GetObservations(f)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}

	for _, d := range data {
		fmt.Println(d.ReferenceTime)
		fmt.Println(d.Observations)
		fmt.Println(d.Geometry)
		fmt.Println(len(d.Observations))
	}

}

// Get location of given station
func getLocation(id string) (lat float64, long float64, err error) {

	f := met.Filter{
		Ids: []string{id},
	}

	stations, err := met.GetSources(f)
	if err != nil {
		return lat, long, err
	}

	if len(stations) != 1 {
		return lat, long, errors.New("Please specify a valid id")
	}

	geom := stations[0].Geometry

	lat = geom.Coordinates[0]
	long = geom.Coordinates[1]

	return lat, long, nil

}
