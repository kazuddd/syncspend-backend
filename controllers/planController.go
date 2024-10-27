package controllers

import (
	"errors"
	"strconv"
	"syncspend/dto"
	"syncspend/helpers"
	"syncspend/services"

	"github.com/gin-gonic/gin"
)

func CreatePlan(c *gin.Context) {
	var request *dto.CreatePlanDTO

	if err := c.ShouldBindJSON(&request); err != nil {
		helpers.ValidationErrorResponse(c, "Invalid request", err.Error())
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		helpers.ErrorResponse(c, errors.New("user not authenticated"))
		return
	}

	request.ID_Owner = strconv.FormatUint(userID.(uint64), 10)

	planResponse, err := (&services.PlanService{}).CreatePlan(*request)
	if err != nil {
		helpers.ErrorResponse(c, err)
		return
	}

	helpers.SuccessResponseWithData(c, "Plan created successfully", planResponse)
}

func UpdatePlan(c *gin.Context) {
	var request *dto.UpdatePlanDTO

	if err := c.ShouldBindJSON(&request); err != nil {
		helpers.ValidationErrorResponse(c, "Invalid request", err.Error())
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		helpers.ErrorResponse(c, errors.New("User not authenticated"))
		return
	}

	request.ID_Owner = strconv.FormatUint(userID.(uint64), 10)

	planID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		helpers.ErrorResponse(c, errors.New("Invalid plan ID"))
		return
	}

	planResponse, err := (&services.PlanService{}).UpdatePlan(planID, *request)
	if err != nil {
		helpers.ErrorResponse(c, err)
		return
	}

	helpers.SuccessResponseWithData(c, "Plan updated successfully", planResponse)
}
