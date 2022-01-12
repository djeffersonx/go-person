package application

import (
	"go-person/internal/infrastructure/bean_factory"
	"net/http"
)

var personService = *bean_factory.PersonService()

func Get(w http.ResponseWriter, r *http.Request) {
	if persons, err := personService.FindAll(); err == nil {
		writeResponse(persons, w)
	} else {
		writeErrorResponse(http.StatusInternalServerError, err, w)
	}
}
