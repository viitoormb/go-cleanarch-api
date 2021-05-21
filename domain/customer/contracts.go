package customer

import (
	"context"
	"time"
)

type CreateCustomer struct {
	Login     string    `json:"login"`
	Document  string    `json:"document"`
	BirthDate time.Time `json:"birthdate"`
	FirstName string    `json:"name"`
	LastName  string    `json:"lastname"`
	Email     string    `json:"email"`
}

type UseCase interface {
	CreateNewCustomer(ctx context.Context, request CreateCustomer) (output *Customer, err error)
	GetAllCustomers(ctx context.Context) ([]*Customer, error)
}

type Repository interface {
	GetAllDocuments(ctx context.Context) ([]*Customer, error)
	CheckIfCustomerExists(ctx context.Context, customer *Customer) (exists bool, err error)
	RegisterCustomer(ctx context.Context, c *Customer) error
}
