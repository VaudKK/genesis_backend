package service

import (
	"context"
	"genesis/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ContributionService interface {
	AddContribution(*models.Contribution) (interface{}, error)
	GetContributionsById(string) (*models.Contribution, error)
}

type contributionService struct {
	collection *mongo.Collection
}

func NewContributionService(c *mongo.Collection) ContributionService {
	return &contributionService{
		collection: c,
	}
}

func (service *contributionService) AddContribution(cntr *models.Contribution) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := service.collection.InsertOne(ctx, cntr)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service *contributionService) GetContributionsById(contributionId string) (*models.Contribution, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var result models.Contribution

	id, err := primitive.ObjectIDFromHex(contributionId)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	filter := bson.D{{Key: "_id", Value: id}}
	err = service.collection.FindOne(ctx, filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		log.Fatal(err)
		return nil, err
	}

	return &result, nil
}
