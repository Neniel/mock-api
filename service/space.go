package service

import (
	"math/rand"

	"github.com/Neniel/mock-api/model"
)

func GetPeopleInSpaceDataMock() *model.Response {
	iss := int32(rand.Intn(10) + 1)
	notIss := int32(rand.Intn(10) + 1)

	response := &model.Response{
		Number:  iss + notIss,
		Message: "success",
	}
	var i int32
	for i = 0; i < iss; i++ {
		response.People = append(response.People, model.Person{Name: "", Craft: "ISS"})
	}

	for i = 0; i < notIss; i++ {
		response.People = append(response.People, model.Person{Name: "", Craft: "dummy"})
	}

	return response
}
