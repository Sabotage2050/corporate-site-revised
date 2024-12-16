package router

import (
	"corporation-site/domain"
	"corporation-site/infra/db/nosql"
	"errors"
	"net/http"
)

type Port int64

type Server interface {
	SetupRouter() error
	Listen() error
	GetHandler() http.Handler
}

var (
	ErrInvalidWebServerInstance = errors.New("invalid web server instance")
)

const (
	InstanceGin int = iota
	InstanceGorillaMux
	InstanceCheckApiGin
)

func NewWebServerFactory(
	instance int,
	port Port,
	dbNoSQL nosql.NoSQLClient,
	emailClient domain.EmailClient,
	logger domain.Logger,
	validator domain.Validator,
) (Server, error) {
	switch instance {
	case InstanceCheckApiGin:
		return NewCheckApiGinServer(port, dbNoSQL, emailClient, logger, validator), nil
	default:
		return nil, ErrInvalidWebServerInstance
	}
}
