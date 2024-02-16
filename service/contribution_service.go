package service

import (
	"context"
	"genesis/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ContributionService interface {
	AddContribution(*models.Contribution) (interface{}, error)
	AddManyContributions([]interface{}) (interface{}, error)
	GetContributionsById(string) (*models.Contribution, error)
	GetContributionsByOrganizationId(string, int, int) ([]models.Contribution, error)
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

func (service *contributionService) AddManyContributions(contributions []interface{}) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := service.collection.InsertMany(ctx, contributions)

	if err != nil {
		return nil, err
	}

	return result.InsertedIDs, nil
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

func (service *contributionService) GetContributionsByOrganizationId(organizationId string, page int, size int) ([]models.Contribution, error) {
	filter := bson.D{{Key: "organizationId", Value: organizationId}}
	opt := options.Find().SetSort(bson.D{{Key: "date", Value: -1}})

	skip := int64(page * size)
	limit := int64(size)
	pagination := options.FindOptions{Limit: &limit, Skip: &skip}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := service.collection.Find(ctx, filter, &pagination, opt)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return []models.Contribution{}, nil
		}

		return nil, err
	}

	result := make([]models.Contribution,0)

	err = cursor.All(ctx, &result)

	if err != nil {
		return nil, err
	}

	return result, nil
}
