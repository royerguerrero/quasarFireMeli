package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"quasarFire/cmd/satellites"

	"github.com/gorilla/mux"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	message := make(map[string]string)
	message["Hello"] = "world"

	satellites := satellites.Satellites{}
	fmt.Println(satellites)
	fmt.Println(satellites.GetMessage())

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}

// Top Secret Service: Calculates the position of the sender of
// the message based on the distance between the satellites and the
// sender and removes noise from the message.
func topSecret(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Top Secret Endpoint")
}

// Top Secret Split Service: Stores the satellite data in
// the hash table [SATELLITES_MEMO] using the client's IP address.
func splitTopSecret(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Top Secret Split Endpoint")
}

// Top Secret Split Read service: Returns the same as the
// /topscret/ service but the satellite data is extracted from the
// table up to [SATELLITES_MEMO].
func splitTopSecretRead(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Top Secret Split Read Endpoint")
}

func main() {
	fmt.Println("Quasar Fire Server")

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", helloWorld)
	router.HandleFunc("/topsecret", topSecret).Methods("POST")
	router.HandleFunc("/topsecret_split", splitTopSecret).Methods("POST")
	router.HandleFunc("/topsecret_split", splitTopSecretRead).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
