package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

type KafkaConsumer struct {
	Client sarama.Consumer
	Topic  string
}

func NewKafkaConsumer(urls []string, topic string) (*KafkaConsumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	conn, err := sarama.NewConsumer(urls, config)
	if err != nil {
		return nil, err
	}

	return &KafkaConsumer{
		Client: conn,
		Topic:  topic,
	}, nil
}

func (k *KafkaConsumer) Consume(ch chan string) error {
	logrus.Infof("Start to consume topic: %s", k.Topic)

	consumer, err := k.Client.ConsumePartition(k.Topic, 0, sarama.OffsetOldest)
	if err != nil {
		return err
	}

	defer consumer.Close()

	messageCh := consumer.Messages()

	for message := range messageCh {
		logrus.Infof("Message received on topic: %s from %s", k.Topic, string(message.Key))
		ch <- string(message.Value)
	}

	return nil
}

func (k *KafkaConsumer) Close() error {
	return k.Client.Close()
}
