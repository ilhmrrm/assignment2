package repository

import "assignment2/model"

type IPersonRepository interface {
	Create(newPerson model.Person) (model.Person, error)
	GetAll() ([]model.Person, error)
	Delete(uuid string) error
}
