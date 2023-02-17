package dao

import (
	"errors"
	"server/models"
)

var peopleList = [...]models.Person{
	{Id: "1", Name: "ale", Age: 28},
	{Id: "2", Name: "mirco", Age: 26},
}

func GetAllPeople() []models.Person {
	return peopleList[:]
}

func GetPersonById(id string) (models.Person, error) {
	for _, person := range peopleList {
		if person.Id == id {
			return person, nil
		}
	}
	return models.Person{}, errors.New("Cannot find id:" + id)
}
