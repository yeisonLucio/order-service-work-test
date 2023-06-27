package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"lucio.com/order-service/src/dto"
	"lucio.com/order-service/src/usecases/contracts"
)

type CustomerController struct {
	CreateCustomerUC contracts.CreateCustomer
}

func (c CustomerController) CreateCustomer(ctx *gin.Context) {
	var createCustomerDTO dto.CreateCustomerDTO

	if err := ctx.ShouldBindJSON(&createCustomerDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"id":    "body_decode_error",
		})
		return
	}

	customer, err := c.CreateCustomerUC.Execute(createCustomerDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
			"id":    "create_customer_error",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": customer,
	})
}
