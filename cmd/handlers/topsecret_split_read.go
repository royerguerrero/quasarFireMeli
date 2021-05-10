package handlers

import (
	"encoding/json"
	"net/http"
	"quasarFire/cmd/data"
)

// Top Secret Split Read service: Returns the same as the
// /topscret/ service but the satellite data is extracted from the
// table up to [SATELLITES_MEMO].
func SplitTopSecretRead(w http.ResponseWriter, r *http.Request) {
	x, y := data.GetLocation(100.0, 115.5, 142.7)
	message := data.GetMessage(
		[]string{"este", "", "", "mensaje", ""},
		[]string{"", "es", "", "", "secreto"},
		[]string{"este", "", "un", "", ""},
	)
	position := make(map[string]float64)
	position["x"] = x
	position["y"] = y

	responsePayload := data.TopSecretResponse{position, message}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responsePayload)
}
