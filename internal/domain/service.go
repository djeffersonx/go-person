package domain

type PersonService struct {
	PersonRepository *PersonRepository
}

func (service PersonService) FindAll() ([]Person, error) {
	return (*service.PersonRepository).FindAll()
}
