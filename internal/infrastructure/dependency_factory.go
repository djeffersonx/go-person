package infrastructure

import (
	"go-person/internal/domain"
	"sync"
)

var (
	personRepository *domain.PersonRepository
	personService    *domain.PersonService
	once             sync.Once
)

func init() {
	once.Do(func() {
		var instance domain.PersonRepository = new(PostgresPersonRepository)
		personRepository = &instance
		personService = &domain.PersonService{PersonRepository: personRepository}
	})
}

func PersonRepository() *domain.PersonRepository {
	return personRepository
}

func PersonService() *domain.PersonService {
	return personService
}
