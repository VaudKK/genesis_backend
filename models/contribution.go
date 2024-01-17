package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Contribution struct {
	Id     primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Name   string             `json:"name,omitempty" validate:"required"`
	Amount float64            `json:"amount,omitempty" validate:"required"`
	Date   int64              `json:"date,omitempty"`
}
