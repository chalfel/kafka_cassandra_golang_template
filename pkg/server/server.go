package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/chalfel/kafka_cassandra_golang_template/pkg/router"
	"github.com/sirupsen/logrus"
)

type Server struct {
	Server *http.Server
	App    *router.App
}

func NewServer(app *router.App, port string) *Server {
	s := &Server{
		App: app,
	}

	s.Server = &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: s.App.Engine,
	}

	return s
}

func (s *Server) Init() error {
	go func() {
		logrus.Infof("Server is up and listening on port: %s", s.Server.Addr)
		if err := s.Server.ListenAndServe(); err != nil {
			logrus.Infoln("Server was shutted down")
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	logrus.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Server.Shutdown(ctx); err != nil {
		logrus.WithError(err).Fatal("Servier forced to shutdown: ")
	}

	logrus.Println("Servier exiting")

	return nil
}
