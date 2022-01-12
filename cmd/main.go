package main

import (
	"github.com/gorilla/mux"
	"go-person/internal/application"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/persons", application.Get)
	log.Fatal(http.ListenAndServe(":8000", router))

}
