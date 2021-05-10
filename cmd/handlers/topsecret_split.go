package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"quasarFire/cmd/data"
	"reflect"

	"github.com/gorilla/mux"
)

// Top Secret Split Service: Stores the satellite data in
// the hash table [SATELLITES_MEMO] using the client's IP address.
func SplitTopSecret(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	satelliteName := vars["name"]

	satellite := data.Satellites.FindByName(satelliteName)

	w.Header().Set("Content-Type", "application/json")

	if reflect.ValueOf(satellite).Field(0).IsZero() {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(data.ErrorResponse{Detail: "Not enough information."})
	} else {
		var newSatellite data.Satellite

		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(data.ErrorResponse{Detail: "Please, ensure to provide a valid satellite."})
		}

		json.Unmarshal(reqBody, &newSatellite)
		// data.Satellites = append(data.Satellites, newSatellite)
	}
}
