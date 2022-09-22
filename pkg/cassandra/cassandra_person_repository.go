package cassandra

import (
	"github.com/chalfel/kafka_cassandra_golang_template/pkg/person"
	"github.com/sirupsen/logrus"
)

type CassandraPersonRepository struct {
	CassandraRepository
	PersonSchema PersonSchema
}

func NewCassandraPersonRepository(c CassandraClient, p PersonSchema) *CassandraPersonRepository {
	return &CassandraPersonRepository{
		CassandraRepository: CassandraRepository{CassandraClient: c},
		PersonSchema:        p,
	}
}

func (c *CassandraPersonRepository) Create(person person.Person) *person.Person {

	q := c.CassandraClient.Session.Query(c.PersonSchema.Table.Insert()).BindStruct(&person)
	if err := q.ExecRelease(); err != nil {
		logrus.Errorf(err.Error())
		return nil
	}

	return &person
}
