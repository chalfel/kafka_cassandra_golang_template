package kafka

import "github.com/sirupsen/logrus"

type KafkaPersonConsumer struct {
	KafkaConsumer KafkaConsumer
}

func NewKafkaPersonConsumer(kafkaConsumer KafkaConsumer) *KafkaPersonConsumer {
	return &KafkaPersonConsumer{

		KafkaConsumer: kafkaConsumer,
	}
}

func (k *KafkaPersonConsumer) Consume() error {
	ch := make(chan string, 1)

	go k.KafkaConsumer.Consume(ch)

	for message := range ch {
		logrus.Infof("Message Received %s", message)
	}

	return nil
}
