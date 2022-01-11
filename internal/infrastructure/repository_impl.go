package infrastructure

import (
	"go-person/internal/domain"
)

type PostgresPersonRepository struct {
	domain.PersonRepository
}

func (PostgresPersonRepository) FindById() (domain.Person, error) {
	return domain.Person{}, nil
}

func (PostgresPersonRepository) FindAll() ([]domain.Person, error) {
	return []domain.Person{
		{ID: "1", Name: "Djefferson william da silva", Email: "djeffersonx@gmail.com"},
	}, nil
}

func (PostgresPersonRepository) FindByEmail(email string) (domain.Person, error) {
	return domain.Person{}, nil
}
