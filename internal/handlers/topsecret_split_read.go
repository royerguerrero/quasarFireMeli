package handlers

import (
	"encoding/json"
	"net/http"
	"quasarFire/internal"
	"reflect"
)

// Top Secret Split Read service: Returns the same as the
// /topscret/ service but the satellite data is extracted from the
// Satellites struct.
func SplitTopSecretRead(w http.ResponseWriter, r *http.Request) {
	satellites := internal.Satellites
	w.Header().Set("Content-Type", "application/json")

	if reflect.ValueOf(satellites).Field(0).IsZero() {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(internal.ErrorResponse{Detail: "Not enough information."})
	} else {
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
