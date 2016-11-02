package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

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

	port := os.Getenv("HTTP_PLATFORM_PORT")

	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(":"+port, router))
}
