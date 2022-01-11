package domain

type PersonRepository interface {
	FindById() (Person, error)
	FindAll() ([]Person, error)
	FindByEmail(email string) (Person, error)
}
