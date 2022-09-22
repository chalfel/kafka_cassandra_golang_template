package cassandra

import "github.com/scylladb/gocqlx/v2/table"

type PersonSchema struct {
	Table table.Table
}

func NewPersonSchema() *PersonSchema {
	var personMetadata = table.Metadata{
		Name:    "person",
		Columns: []string{"first_name", "last_name", "email"},
		PartKey: []string{"first_name"},
		SortKey: []string{"last_name"},
	}

	var personTable = table.New(personMetadata)

	return &PersonSchema{
		Table: *personTable,
	}
}
