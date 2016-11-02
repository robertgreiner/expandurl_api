package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getExpandedURLEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	response := ExpandURL(params["name"])

	json.NewEncoder(w).Encode(response)
}

func main() {
	router := mux.NewRouter().UseEncodedPath()
	router.HandleFunc("/expand/{name}", getExpandedURLEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
