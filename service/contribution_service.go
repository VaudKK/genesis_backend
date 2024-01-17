package service

import (
	"context"
	"genesis/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type ContributionService interface {
	AddContribution(*models.Contribution) (interface{}, error)
	//GetContributionsByGroup(int) []models.Contribution
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

// func (service *contributionService) GetContributionsByGroup(groupId int) []models.Contribution {

// }
