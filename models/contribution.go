package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Contribution struct {
	Id          primitive.ObjectID  `bson:"_id" json:"_id,omitempty"`
	Name        string              `json:"name,omitempty" validate:"required" binding:"max=30"`
	Phone       string              `json:"phone" validate:"required" binding:"max=30"`
	Amount      float64             `json:"amount,omitempty" validate:"required" binding:"gte=1,lte=100000000"`
	Date        int64               `json:"date,omitempty" validate:"required"`
	PaymentMode string              `json:"paymentMode" validate:"required" binding:"min=4,max=15"`
	CreatedAt   primitive.Timestamp `json:"createdAt"`
}

type ContributionList struct {
	Contributions []*Contribution `json:"contributions" validate:"dive"`
}
