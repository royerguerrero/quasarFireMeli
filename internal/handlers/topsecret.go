package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"quasarFire/internal"
)

// Top Secret Service: Calculates the position of the sender of
// the message based on the distance between the satellites and the
// sender and removes noise from the message.
func TopSecret(w http.ResponseWriter, r *http.Request) {
	var newSatellites internal.SatellitesRepo
	w.Header().Set("Content-Type", "application/json")

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Please, ensure to provide the satellites")
	}

	json.Unmarshal(reqBody, &newSatellites)
	err = newSatellites.Validate()

	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(internal.ErrorResponse{Detail: err})
	} else {
		internal.Satellites = newSatellites
		distances := []float64{}
		messages := [][]string{}

		for _, s := range internal.Satellites.Satellites {
			distances = append(distances, s.Distance)
			messages = append(messages, s.Message)
		}

		x, y := internal.GetLocation(distances...)
		message := internal.GetMessage(messages...)
		position := make(map[string]float64)
		position["x"] = x
		position["y"] = y

		responsePayload := internal.TopSecretResponse{Position: position, Message: message}

		json.NewEncoder(w).Encode(responsePayload)
	}
}
