package router

import (
	"github.com/chalfel/kafka_cassandra_golang_template/pkg/person"
	"github.com/gin-gonic/gin"
)

func (a *App) CreatePerson(c *gin.Context) {
	p := person.Person{
		FirstName: "Caio",
		LastName:  "Felix",
		Email:     []string{"caiohalcsik@gmail.com"},
	}

	newP := a.PersonRepository.Create(p)

	a.PersonProducer.SendMessage("teste")

	c.JSON(200, gin.H{"data": newP})
}
