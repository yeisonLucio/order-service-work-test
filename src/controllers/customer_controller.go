package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"lucio.com/order-service/src/dto"
	rcontracts "lucio.com/order-service/src/repositories/contracts"
	"lucio.com/order-service/src/usecases/contracts"
)

type CustomerController struct {
	CreateCustomerUC    contracts.CreateCustomerUC
	CreateWorkOrderUC   contracts.CreateWorkOrderUC
	CustomerRepository  rcontracts.CustomerRepository
	WorkOrderRepository rcontracts.WorkOrderRepository
	UpdateCustomerUC    contracts.UpdateCustomerUC
}

func (c CustomerController) CreateCustomer(ctx *gin.Context) {
	var createCustomerDTO dto.CreateCustomerDTO

	if err := ctx.ShouldBindJSON(&createCustomerDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"id":    "bad_request",
		})
		return
	}

	customer, err := c.CreateCustomerUC.Execute(createCustomerDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
			"id":    "unexpected_error",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": customer,
	})
}

func (c *CustomerController) CreateWorkOrder(ctx *gin.Context) {
	var createWorkOrderDTO dto.CreateWorkOrderDTO

	if err := ctx.ShouldBindJSON(&createWorkOrderDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"id":    "bad_request",
		})
		return
	}

	createWorkOrderDTO.CustomerID = ctx.Param("id")

	workOrder, err := c.CreateWorkOrderUC.Execute(createWorkOrderDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
			"id":    "unexpected_error",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": workOrder,
	})
}

func (w *CustomerController) GetWorkOrders(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": w.WorkOrderRepository.GetByCustomerID(ctx.Param("id")),
	})
}

func (w *CustomerController) GetCustomers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": w.CustomerRepository.GetActives(),
	})
}

func (w *CustomerController) GetCustomer(ctx *gin.Context) {
	customer, err := w.CustomerRepository.FindByID(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
			"id":    "record_not_found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": customer,
	})
}

func (w *CustomerController) UpdateCustomer(ctx *gin.Context) {
	var updateCustomerDTO dto.UpdateCustomerDTO

	if err := ctx.ShouldBindJSON(&updateCustomerDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"id":    "bad_request",
		})
		return
	}

	updateCustomerDTO.ID = ctx.Param("id")

	updatedCustomer, err := w.UpdateCustomerUC.Execute(updateCustomerDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
			"id":    "unexpected_error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": updatedCustomer,
	})
}

func (w *CustomerController) DeleteCustomer(ctx *gin.Context) {
	if err := w.CustomerRepository.DeleteByID(ctx.Param("id")); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
			"id":    "record_not_found",
		})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
