package kafka

type KafkaPersonProducer struct {
	KafkaProducer KafkaProducer
}

func NewKafkaPersonProducer(k KafkaProducer) *KafkaPersonProducer {
	return &KafkaPersonProducer{
		KafkaProducer: k,
	}
}

func (k *KafkaPersonProducer) SendMessage(message string) error {
	return k.KafkaProducer.SendMessage(message)
}
