package domain

type Validator interface {
	Validate(interface{}) error
	Messages() []string
}
