package router

import (
	"github.com/chalfel/kafka_cassandra_golang_template/pkg/config"
	"github.com/chalfel/kafka_cassandra_golang_template/pkg/middleware"
	"github.com/chalfel/kafka_cassandra_golang_template/pkg/person"
	"github.com/gin-gonic/gin"
)

type App struct {
	Engine           *gin.Engine
	Cfg              config.AppConfig
	PersonRepository person.PersonRepository
	PersonProducer   person.PersonProducer
}

func NewApp(
	cfg config.AppConfig,
	personRepository person.PersonRepository,
	personProducer person.PersonProducer,
) (*App, error) {
	var err error

	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	r := App{
		Engine:           router,
		Cfg:              cfg,
		PersonRepository: personRepository,
		PersonProducer:   personProducer,
	}

	return &r, err
}

func (a *App) RegisterRoutes() {
	a.Engine.Use(middleware.ErrorHandler())
	a.Engine.GET("/health/check", a.HealthCheck)
	a.Engine.POST("/person", a.CreatePerson)
}
