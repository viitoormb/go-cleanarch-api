package repository

import (
	"context"
	"log"

	"github.com/viitoormb/go-cleanarch-api/domain/customer"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type customerMongoRepository struct {
	db *mongo.Database
}

func NewCustomerMongoRepository(db *mongo.Client) *customerMongoRepository {
	return &customerMongoRepository{
		db: db.Database("service"),
	}
}

func (r *customerMongoRepository) GetAllDocuments(ctx context.Context) ([]*customer.Customer, error) {
	result, err := r.db.Collection("customer").Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	var customers []*customer.Customer

	for result.Next(ctx) {
		var customer *customer.Customer

		err := result.Decode(&customer)
		if err != nil {
			log.Fatal(err)
		}
		customers = append(customers, customer)
	}

	return customers, nil
}

func (repo *customerMongoRepository) RegisterCustomer(ctx context.Context, c *customer.Customer) error {
	result, err := repo.db.Collection("customer").InsertOne(ctx, c)
	if err != nil {
		return err
	}
	log.Println("Inserido", result)
	return nil
}

func (r *customerMongoRepository) CheckIfCustomerExists(ctx context.Context, c *customer.Customer) (bool, error) {
	var customers []*customer.Customer
	c.ID = bson.NewObjectId()

	result, err := r.db.Collection("customer").Find(ctx, bson.M{"document": c.Document})

	if err != nil {
		return false, err
	}

	for result.Next(ctx) {
		var customer *customer.Customer

		err := result.Decode(&c)
		if err != nil {
			log.Fatal(err)
		}

		customers = append(customers, customer)
	}

	if len(customers) > 0 {
		return true, nil
	}

	return false, nil
}
