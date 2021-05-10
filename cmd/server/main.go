package main

import (
	"fmt"
	"log"
	"net/http"
	"quasarFire/internal/handlers"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println(`
	________                                      ___________.__                
	\_____  \  __ _______    ___________ _______  \_   _____/|__|______   ____  
	 /  / \  \|  |  \__  \  /  ___/\__  \\_  __ \  |    __)  |  \_  __ \_/ __ \ 
	/   \_/.  \  |  // __ \_\___ \  / __ \|  | \/  |     \   |  ||  | \/\  ___/ 
	\_____\ \_/____/(____  /____  >(____  /__|     \___  /   |__||__|    \___  >
		\__>          \/     \/      \/             \/                    \/ 

				Made with 💛 by @royerguerrero 
	`)

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/topsecret/", handlers.TopSecret).Methods("POST")
	router.HandleFunc("/topsecret_split/{name:[A-Za-z]+}/", handlers.SplitTopSecret).Methods("POST")
	router.HandleFunc("/topsecret_split/", handlers.SplitTopSecretRead).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
