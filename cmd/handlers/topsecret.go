package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"quasarFire/cmd/data"
)

// Top Secret Service: Calculates the position of the sender of
// the message based on the distance between the satellites and the
// sender and removes noise from the message.
func TopSecret(w http.ResponseWriter, r *http.Request) {
	var newSatellites data.SatellitesRepo

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Please, ensure to provide the satellites")
	}

	json.Unmarshal(reqBody, &newSatellites)

	data.Satellites = newSatellites
	distances := []float64{}
	messages := [][]string{}

	for _, s := range data.Satellites.Satellites {
		distances = append(distances, s.Distance)
		messages = append(messages, s.Message)
	}

	x, y := data.GetLocation(distances...)
	message := data.GetMessage(messages...)
	position := make(map[string]float64)
	position["x"] = x
	position["y"] = y

	fmt.Println(position, message)

	responsePayload := data.TopSecretResponse{Position: position, Message: message}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responsePayload)
}
