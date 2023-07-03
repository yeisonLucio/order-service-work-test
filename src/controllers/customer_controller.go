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

// @Summary Servicio para crear clientes
// @Description Permite crear un determinado cliente
// @Tags Customers
// @Accept json
// @Produce json
// @Param body body dto.CreateCustomerDTO true "Body data"
// @Success 201 {object} dto.CreatedCustomerDTO
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /customers [post]
func (c CustomerController) CreateCustomer(ctx *gin.Context) {
	var createCustomerDTO dto.CreateCustomerDTO

	if err := ctx.ShouldBindJSON(&createCustomerDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			ID:      "bad_request",
			Message: err.Error(),
		})
		return
	}

	customer, err := c.CreateCustomerUC.Execute(createCustomerDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			ID:      "unexpected_error",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, customer)
}

// @Summary Servicio para crear una orden de servicio para un cliente
// @Description Permite crear una orden de servicio para un cliente
// @Tags Customers
// @Accept json
// @Produce json
// @param id path string true "id del cliente"
// @Param body body dto.CreateWorkOrderDTO true "Body data"
// @Success 201 {object} dto.CreatedWorkOrderDTO
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /customers/{id}/work-orders [post]
func (c *CustomerController) CreateWorkOrder(ctx *gin.Context) {
	var createWorkOrderDTO dto.CreateWorkOrderDTO

	if err := ctx.ShouldBindJSON(&createWorkOrderDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			ID:      "bad_request",
			Message: err.Error(),
		})
		return
	}

	createWorkOrderDTO.CustomerID = ctx.Param("id")

	workOrder, err := c.CreateWorkOrderUC.Execute(createWorkOrderDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			ID:      "unexpected_error",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, workOrder)
}

// @Summary Servicio para obtener las ordenes de servicio de un cliente
// @Description Este servicio permite obtener todas las ordenes de servicio de un cliente
// @Tags Customers
// @Accept json
// @Produce json
// @param id path string true "id del cliente"
// @Success 200 {array} dto.WorkOrderDTO
// @Router /customers/{id}/work-orders [get]
func (w *CustomerController) GetWorkOrders(ctx *gin.Context) {
	workOrders := w.WorkOrderRepository.GetByCustomerID(ctx.Param("id"))
	ctx.JSON(http.StatusOK, workOrders)
}

// @Summary Servicio para obtener todos los clientes activos
// @Description Este servicio permite obtener todos los clientes activos
// @Tags Customers
// @Accept json
// @Produce json
// @Success 200 {array} dto.CustomerDTO
// @Router /customers [get]
func (w *CustomerController) GetCustomers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, w.CustomerRepository.GetActives())
}

// @Summary Servicio para obtener un cliente
// @Description Este servicio permite obtener un cliente
// @Tags Customers
// @Accept json
// @Produce json
// @param id path string true "id del cliente"
// @Success 200 {object} dto.CustomerDTO
// @Failure 404 {object} dto.ErrorResponse
// @Router /customers/{id} [get]
func (w *CustomerController) GetCustomer(ctx *gin.Context) {
	customer, err := w.CustomerRepository.GetByID(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.ErrorResponse{
			ID:      "record_not_found",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, customer)
}

// @Summary Servicio para actualizar un cliente
// @Description Este servicio permite realizar la actualización de un cliente
// @Tags Customers
// @Accept json
// @Produce json
// @param id path string true "id del cliente"
// @Param body body dto.UpdateCustomerDTO true "Body data"
// @Success 200 {object} dto.CustomerDTO
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /customers/{id} [put]
func (w *CustomerController) UpdateCustomer(ctx *gin.Context) {
	var updateCustomerDTO dto.UpdateCustomerDTO

	if err := ctx.ShouldBindJSON(&updateCustomerDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			ID:      "bad_request",
			Message: err.Error(),
		})
		return
	}

	updateCustomerDTO.ID = ctx.Param("id")

	updatedCustomer, err := w.UpdateCustomerUC.Execute(updateCustomerDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			ID:      "unexpected_error",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, updatedCustomer)
}

// @Summary Servicio para eliminar un cliente
// @Description Este servicio permite realizar la eliminación de un cliente
// @Tags Customers
// @Accept json
// @Produce json
// @param id path string true "id del cliente"
// @Success 204 "No Content"
// @Failure 404 {object} dto.ErrorResponse
// @Router /customers/{id} [delete]
func (w *CustomerController) DeleteCustomer(ctx *gin.Context) {
	if err := w.CustomerRepository.DeleteByID(ctx.Param("id")); err != nil {
		ctx.JSON(http.StatusNotFound, dto.ErrorResponse{
			ID:      "record_not_found",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
