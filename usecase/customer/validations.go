package usecase

import (
	"context"
	"fmt"

	"github.com/viitoormb/go-cleanarch-api/domain/customer"
)

func checkIfUserCanBeCreated(ctx context.Context, c *customer.Customer, repository customer.Repository) error {

	if c.Age() < 18 {
		return customer.CustomerUnderAgeError()
	}

	userExits, err := repository.CheckIfCustomerExists(ctx, c)

	if err != nil {
		fmt.Errorf("error on check if user exists", err)
		return err
	}

	if userExits {
		return customer.CustomerAlreadyRegisteredError()
	}

	return nil
}
