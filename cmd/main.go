package main

import (
	"github.com/gorilla/mux"
	"go-person/internal/application"
	"go-person/pkg/monitoring"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/persons", application.Get).Methods("GET")
	router.HandleFunc("/persons", application.Post).Methods("POST")

	monitoring.Init(router)

	log.Fatal(http.ListenAndServe(":8000", router))

}
