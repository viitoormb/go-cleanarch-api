package customer

import (
	"errors"

	"github.com/viitoormb/go-cleanarch-api/domain/shared"
)

func CustomerUnderAgeError() error {
	return &shared.DomainError{
		Err:        errors.New("use has not age to register"),
		StatusCode: 422,
	}
}

func CustomerAlreadyRegisteredError() error {
	return &shared.DomainError{
		Err:        errors.New("user already registered"),
		StatusCode: 422,
	}
}
