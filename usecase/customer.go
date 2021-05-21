package usecase

import (
	"context"

	"github.com/viitoormb/go-cleanarch-api/domain/customer"
)

type customerUseCase struct {
	repository customer.Repository
}

func NewCustomerUseCase(repository customer.Repository) *customerUseCase {
	return &customerUseCase{
		repository: repository,
	}
}

func (useCase *customerUseCase) CreateNewCustomer(ctx context.Context, request customer.CreateCustomer) (c *customer.Customer, err error) {
	c = &customer.Customer{
		Login:     request.Login,
		Document:  request.Document,
		BirthDate: request.BirthDate,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
	}

	if err := checkIfUserCanBeCreated(ctx, c, useCase.repository); err != nil {
		return nil, err
	}

	err = useCase.repository.RegisterCustomer(ctx, c)
	return c, err
}

func (useCase *customerUseCase) GetAllCustomers(ctx context.Context) (output []*customer.Customer, err error) {
	output, err = useCase.repository.GetAllDocuments(ctx)
	return output, err
}
