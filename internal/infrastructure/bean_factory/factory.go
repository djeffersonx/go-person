package bean_factory

import (
	"go-person/internal/domain"
	"go-person/internal/infrastructure"
	"sync"
)

// todo remove bean management, have more of functional mindset !!!

var (
	once sync.Once
)

func init() {
	once.Do(func() {
		createPersonRepository()
		createPersonService(personRepository)
	})
}

func createPersonService(personRepository *domain.PersonRepository) {
	personService = &domain.PersonService{PersonRepository: personRepository}
}

func createPersonRepository() {
	var instance domain.PersonRepository = new(infrastructure.PostgresPersonRepository)
	personRepository = &instance
}
