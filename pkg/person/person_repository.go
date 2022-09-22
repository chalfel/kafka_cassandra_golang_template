package person

type PersonRepository interface {
	Create(person Person) *Person
}
