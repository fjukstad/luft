package controllers

import (
	"net/http"

	"github.com/fjukstad/met"
)

func PrecipitationHandler(w http.ResponseWriter, r *http.Request) {
	f := met.Filter{
		Sources:       []string{"SN90450"},
		ReferenceTime: "2017-01-01T00:00:00.000Z/2017-03-2T00:00:00.000Z",
		Elements:      []string{"surface_snow_thickness"},
	}

	_, err := met.GetObservations(f)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
