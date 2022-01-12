package bean_factory

import (
	"go-person/internal/domain"
)

var (
	personRepository *domain.PersonRepository
	personService    *domain.PersonService
)

func PersonRepository() *domain.PersonRepository {
	return personRepository
}

func PersonService() *domain.PersonService {
	return personService
}
