package person

type PersonProducer interface {
	SendMessage(message string) error
}
