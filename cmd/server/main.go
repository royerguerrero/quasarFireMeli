package main

import (
	"fmt"
	"log"
	"net/http"
	"quasarFire/cmd/handlers"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Quasar Fire Server")

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/topsecret/", handlers.TopSecret).Methods("POST")
	router.HandleFunc("/topsecret_split/{name:[A-Za-z]+}/", handlers.SplitTopSecret).Methods("POST")
	router.HandleFunc("/topsecret_split/", handlers.SplitTopSecretRead).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
