package log

import (
	"corporation-site/domain"
	"errors"
)

const (
	InstanceLogrusLogger int = iota
	InstanceZapLogger
)

var (
	errInvalidLoggerInstance = errors.New("invalid log instance")
)

func NewLoggerFactory(instance int) (domain.Logger, error) {
	switch instance {
	case InstanceZapLogger:
		return NewZapLogger()
	case InstanceLogrusLogger:
		return NewLogrusLogger(), nil
	default:
		return nil, errInvalidLoggerInstance
	}
}
