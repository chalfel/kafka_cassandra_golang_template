package main

import (
	"github.com/chalfel/kafka_cassandra_golang_template/cmd"
	"github.com/sirupsen/logrus"
)

func main() {
	root := cmd.NewRootCmd()
	if err := root.Execute(); err != nil {
		logrus.WithError(err).Fatal("command resulted in error")
	}
}
