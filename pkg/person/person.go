package person

type Person struct {
	FirstName string
	LastName  string
	Email     []string
	HairColor string `db:"-"`
}
