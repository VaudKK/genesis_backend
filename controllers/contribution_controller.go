package controllers

import (
	"genesis/models"
	"genesis/responses"
	"genesis/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var validate = validator.New()

type ContributionController interface {
	AddContribution() gin.HandlerFunc
}

type contributionController struct {
	service service.ContributionService
}

func NewContributionController(svc service.ContributionService) ContributionController {
	return &contributionController{
		service: svc,
	}
}

func (controller contributionController) AddContribution() gin.HandlerFunc {
	return func(c *gin.Context) {
		var Contribution models.Contribution

		if err := c.BindJSON(&Contribution); err != nil {
			c.JSON(http.StatusBadRequest, responses.ContributionResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
		}

		//check required fields
		if validatorErr := validate.Struct(&Contribution); validatorErr != nil {
			c.JSON(http.StatusBadRequest, responses.ContributionResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    map[string]interface{}{"data": validatorErr.Error()},
			})
			return
		}

		newContribution := models.Contribution{
			Id:     primitive.NewObjectID(),
			Name:   Contribution.Name,
			Amount: Contribution.Amount,
			Date:   time.Now().Unix(),
		}

		result, addErr := controller.service.AddContribution(&newContribution)

		if addErr != nil {
			c.JSON(http.StatusBadRequest, responses.ContributionResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    map[string]interface{}{"data": addErr.Error()},
			})
			return
		}

		c.JSON(http.StatusCreated, responses.ContributionResponse{
			Status:  http.StatusCreated,
			Message: "Contribution Added",
			Data:    map[string]interface{}{"data": result},
		})
	}
}
