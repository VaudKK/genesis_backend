package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Contribution struct {
	Id     primitive.ObjectID `json:"id,omitempty"`
	Name   string             `json:"name,omitempty" validate:"required"`
	Amount float64            `json:"amount,omitempty" validate:"required"`
	Date   primitive.DateTime `json:"date,omitempty" validate:"required"`
}
