package server

import (
	"corporation-site/domain"
	"corporation-site/infra/db/nosql"
	emailinfra "corporation-site/infra/email"
	"corporation-site/infra/log"
	"corporation-site/infra/router"
	"corporation-site/infra/validation"

	"strconv"
)

type Config struct {
	appName       string
	dbNoSQL       nosql.NoSQLClient
	emailClient   domain.EmailClient
	webServerPort router.Port
	webServer     router.Server
	logger        domain.Logger
	validator     domain.Validator
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Name(name string) *Config {
	c.appName = name
	return c
}

func (c *Config) WebServerPort(port string) *Config {
	p, err := strconv.ParseInt(port, 10, 64)
	if err != nil {
		panic(err)
	}
	c.webServerPort = router.Port(p)
	return c
}

func (c *Config) DbNoSQL(instance int) *Config {
	dbNoSQL, err := nosql.NewNosqlDBFactory(nosql.DynamoDBStorage)
	if err != nil {
		return nil
	}
	c.dbNoSQL = dbNoSQL
	return c
}

func (c *Config) Email(provider int) *Config {

	emailClient, err := emailinfra.NewEmailClientFactory(provider)
	if err != nil {
		return nil
	}
	c.emailClient = emailClient
	return c
}

func (c *Config) WebServer(instance int) *Config {
	s, err := router.NewWebServerFactory(
		instance,
		c.webServerPort,
		c.dbNoSQL,
		c.emailClient,
		c.logger,
		c.validator,
	)
	if err != nil {
		panic(err)
	}
	c.webServer = s
	return c
}

func (c *Config) Start() error {
	c.webServer.Listen()
	return nil
}

func (c *Config) Logger(instance int) *Config {
	logger, err := log.NewLoggerFactory(instance)
	if err != nil {
		panic(err)
	}
	c.logger = logger
	return c
}

func (c *Config) Validator(instance int) *Config {
	validator, err := validation.NewValidatorFactory(instance)
	if err != nil {
		panic(err)
	}
	c.validator = validator
	return c
}
