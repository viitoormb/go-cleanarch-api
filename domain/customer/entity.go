package customer

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Customer struct {
	ID        bson.ObjectId `bson:"_id" json:"_id,omitempty"`
	Login     string        `json:"login"`
	Document  string        `json:"document"`
	BirthDate time.Time     `json:"birthdate"`
	FirstName string        `json:"name"`
	LastName  string        `json:"lastname"`
	Email     string        `json:"email"`
}

func (c *Customer) Age() int {
	years := time.Now().Year() - c.BirthDate.Year()
	return years
}
