package main

import (
	"context"
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/viitoormb/go-cleanarch-api/api/controllers"
	"github.com/viitoormb/go-cleanarch-api/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Errorf("Error on connect on db %v", err)
		panic(err)
	}
	fmt.Printf("%+v\n", client)
	customerRepo := repository.NewCustomerMongoRepository(client)

	controllers.NewCustomerController(e, customerRepo)

	e.Start(":3009")
}
