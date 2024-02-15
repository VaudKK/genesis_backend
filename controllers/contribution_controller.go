package controllers

import (
	"genesis/models"
	"genesis/responses"
	"genesis/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var validate = validator.New()

type ContributionController interface {
	AddContribution() gin.HandlerFunc
	AddManyContributions() gin.HandlerFunc
	GetContributionById() gin.HandlerFunc
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
				Data:    err.Error(),
			})
			c.Abort()
			return
		}

		//check required fields
		if validatorErr := validate.Struct(&Contribution); validatorErr != nil {
			c.JSON(http.StatusBadRequest, responses.ContributionResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    validatorErr.Error(),
			})
			c.Abort()
			return
		}

		newContribution := models.Contribution{
			Id:          primitive.NewObjectID(),
			Name:        Contribution.Name,
			Phone:		 Contribution.Phone,
			Amount:      Contribution.Amount,
			Date:        Contribution.Date,
			PaymentMode: Contribution.PaymentMode,
		}

		result, addErr := controller.service.AddContribution(&newContribution)

		if addErr != nil {
			c.JSON(http.StatusBadRequest, responses.ContributionResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    addErr.Error(),
			})
			c.Abort()
			return
		}

		c.JSON(http.StatusCreated, responses.ContributionResponse{
			Status:  http.StatusCreated,
			Message: "Contribution Added",
			Data:    result,
		})
	}
}

func (controller contributionController) AddManyContributions() gin.HandlerFunc {
	return func(c *gin.Context) {
		var contributions models.ContributionList

		if err := c.BindJSON(&contributions); err != nil {
			c.JSON(http.StatusBadRequest, responses.ContributionResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    err.Error(),
			})
			c.Abort()
			return
		}

		//check required fields
		if validatorErr := validate.Struct(&contributions); validatorErr != nil {
			c.JSON(http.StatusBadRequest, responses.ContributionResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    validatorErr.Error(),
			})
			c.Abort()
			return
		}

		newContributions := buildContributions(contributions)

		result, addErr := controller.service.AddManyContributions(newContributions)

		if addErr != nil {
			c.JSON(http.StatusBadRequest, responses.ContributionResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    addErr.Error(),
			})
			c.Abort()
			return
		}

		c.JSON(http.StatusCreated, responses.ContributionResponse{
			Status:  http.StatusCreated,
			Message: "Contributions Added",
			Data:    result,
		})
	}
}

func buildContributions(contributions models.ContributionList) []interface{} {
	newContributions := make([]interface{}, 0)
	for _, contribution := range contributions.Contributions {
		newContributions = append(newContributions, models.Contribution{
			Id:          primitive.NewObjectID(),
			Name:        contribution.Name,
			Phone:		 contribution.Phone,
			Amount:      contribution.Amount,
			Date:        contribution.Date,
			PaymentMode: contribution.PaymentMode,
		})
	}

	return newContributions
}

func (controller contributionController) GetContributionById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		if id == "" {
			c.JSON(http.StatusBadRequest, responses.ContributionResponse{
				Status:  http.StatusBadRequest,
				Message: "Missing path parameter",
				Data:    "Missing path parameter",
			})

			return
		}

		result, err := controller.service.GetContributionsById(id)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ContributionResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    err,
			})
			return
		}

		c.JSON(http.StatusOK, responses.ContributionResponse{
			Status:  http.StatusOK,
			Message: "Success",
			Data:    result,
		})
	}
}
