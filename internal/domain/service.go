package domain

import (
	"go-person/pkg/monitoring"
)

type PersonService struct {
	PersonRepository *PersonRepository
}

func (service PersonService) FindAll() ([]Person, error) {
	return (*service.PersonRepository).FindAll()
}

func (service PersonService) Create() {
	monitoring.TotalPersonsCreated.Inc()
}
