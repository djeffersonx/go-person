package application

import (
	"go-person/internal/infrastructure"
	"net/http"
)

var personService = *infrastructure.PersonService()

func Get(w http.ResponseWriter, r *http.Request) {
	if persons, err := personService.FindAll(); err == nil {
		writeResponse(persons, w)
	} else {
		writeErrorResponse(http.StatusInternalServerError, err, w)
	}
}
