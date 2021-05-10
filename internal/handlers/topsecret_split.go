package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"quasarFire/internal"
	"reflect"

	"github.com/gorilla/mux"
)

// Top Secret Split Service: Stores the satellite data in
// the Satellites struct.
func SplitTopSecret(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	satelliteName := vars["name"]

	satellite, i := internal.Satellites.FindByName(satelliteName)

	w.Header().Set("Content-Type", "application/json")

	if reflect.ValueOf(satellite).Field(0).IsZero() {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(internal.ErrorResponse{Detail: "Not enough information."})
	} else {
		var newSatellite internal.Satellite

		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(internal.ErrorResponse{Detail: "Please, ensure to provide a valid satellite."})
		}

		newSatellite.Name = satellite.Name
		json.Unmarshal(reqBody, &newSatellite)
		err = newSatellite.Validate()

		if err != nil {
			w.WriteHeader(http.StatusNotAcceptable)
			json.NewEncoder(w).Encode(internal.ErrorResponse{Detail: err})
		} else {
			internal.Satellites.Satellites[i] = newSatellite
			SplitTopSecretRead(w, r)
		}
	}
}
