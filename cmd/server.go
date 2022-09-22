package cmd

import (
	"context"

	"github.com/chalfel/kafka_cassandra_golang_template/pkg/cassandra"
	"github.com/chalfel/kafka_cassandra_golang_template/pkg/config"
	"github.com/chalfel/kafka_cassandra_golang_template/pkg/kafka"
	"github.com/chalfel/kafka_cassandra_golang_template/pkg/router"
	"github.com/chalfel/kafka_cassandra_golang_template/pkg/server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func NewServerCmd(ctx context.Context) *cobra.Command {
	command := &cobra.Command{
		Use:   "serve",
		Short: "Serve HTTP application",
		RunE: func(cmd *cobra.Command, args []string) error {
			return StartServer(cmd, args)
		},
	}

	return command
}

func StartServer(cmd *cobra.Command, arg []string) error {
	cfg := &config.AppConfig{}

	if err := config.NewConfig(cfg); err != nil {
		return err
	}

	cassandraClient, err := cassandra.NewCassandraClient([]string{cfg.ScyllaHost}, cfg.PersonKeyspace)

	if err != nil {
		return err
	}

	personSchema := cassandra.NewPersonSchema()
	personRepository := cassandra.NewCassandraPersonRepository(*cassandraClient, *personSchema)

	logrus.Info(cfg.KafkaHost[0])
	kafkaConsumer, err := kafka.NewKafkaConsumer(cfg.KafkaHost, "person")
	kafkaProducer, err := kafka.NewKafkaProducer(cfg.KafkaHost, "person")

	if err != nil {
		return err
	}

	pk := kafka.NewKafkaPersonConsumer(*kafkaConsumer)
	pkp := kafka.NewKafkaPersonProducer(*kafkaProducer)

	go pk.Consume()

	app, err := router.NewApp(*cfg, personRepository, pkp)

	if err != nil {
		return err
	}

	server := server.NewServer(app, cfg.HttpPort)

	app.RegisterRoutes()

	if err := server.Init(); err != nil {
		return err
	}

	return nil
}
