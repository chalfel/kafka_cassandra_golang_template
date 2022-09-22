package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

type KafkaProducer struct {
	Client sarama.SyncProducer
	Topic  string
}

func NewKafkaProducer(urls []string, topic string) (*KafkaProducer, error) {
	config := sarama.NewConfig()

	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	client, err := sarama.NewSyncProducer(urls, config)

	if err != nil {
		return nil, err
	}

	return &KafkaProducer{
		Client: client,
		Topic:  topic,
	}, nil
}

func (k *KafkaProducer) SendMessage(message string) error {
	logrus.Infof("Trying to send message in topic: %s", k.Topic)

	msg := &sarama.ProducerMessage{
		Topic: k.Topic,
		Value: sarama.StringEncoder(message),
	}
	partition, offset, err := k.Client.SendMessage(msg)

	logrus.Infof("Kafka Producer - Message is stored in topic(%s)/partition(%d)/offset(%d)\n", k.Topic, partition, offset)

	return err
}

func (k *KafkaProducer) SendMessages(messages []string) error {
	logrus.Infof("Trying to send messages in topic: %s", k.Topic)

	var msgs []*sarama.ProducerMessage

	for _, message := range messages {
		msg := &sarama.ProducerMessage{
			Topic: k.Topic,
			Value: sarama.StringEncoder(message),
		}

		msgs = append(msgs, msg)
	}

	err := k.Client.SendMessages(msgs)

	logrus.Infof("Kafka Producer - All %d Messages is stored in topic(%s)\n", len(msgs), k.Topic)

	return err
}

func (k *KafkaProducer) Close() error {
	return k.Client.Close()
}
