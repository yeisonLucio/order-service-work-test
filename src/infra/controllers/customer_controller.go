package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"lucio.com/order-service/src/domain/customer/repositories"
	ucConsumer "lucio.com/order-service/src/domain/customer/usecases"
	"lucio.com/order-service/src/domain/workorder/dtos"
	reposWorkOrder "lucio.com/order-service/src/domain/workorder/repositories"
	ucCreateWork "lucio.com/order-service/src/domain/workorder/usecases"
	"lucio.com/order-service/src/infra/requests/customer"
	"lucio.com/order-service/src/infra/requests/workorder"
)

// CustomerController estructura que permite gestionar los clientes
type CustomerController struct {
	CreateCustomerUC    ucConsumer.CreateCustomerUC
	CreateWorkOrderUC   ucCreateWork.CreateWorkOrderUC
	CustomerRepository  repositories.CustomerRepository
	WorkOrderRepository reposWorkOrder.WorkOrderRepository
	UpdateCustomerUC    ucConsumer.UpdateCustomerUC
	Logger              *logrus.Logger
}

// @Summary Servicio para crear clientes
// @Description Permite crear un determinado cliente
// @Tags Customers
// @Accept json
// @Produce json
// @Param body body dto.CreateCustomerDTO true "Body data"
// @Success 201 {object} dto.CreatedCustomerResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /customers [post]
func (c *CustomerController) CreateCustomer(ctx *gin.Context) {
	log := c.Logger.WithFields(logrus.Fields{
		"file":    "customer_controller",
		"method":  "CreateCustomer",
		"request": ctx.Request,
	})

	var createCustomerDTO customer.CreateCustomer

	customerEntity, err := createCustomerDTO.ValidateAndGetEntity(ctx)
	if err != nil {
		log = log.WithField("error", err)
		log.Error()
		ctx.JSON(err.Code, gin.H{
			"error": err.Error.Error(),
		})
		return
	}

	log = log.WithField("customer", customerEntity)

	customer, err := c.CreateCustomerUC.Execute(*customerEntity)
	if err != nil {
		log = log.WithField("error", err)
		log.Error()
		ctx.JSON(err.Code, gin.H{
			"error": err.Error.Error(),
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
	log := c.Logger.WithFields(logrus.Fields{
		"file":    "customer_controller",
		"method":  "CreateWorkOrder",
		"request": ctx.Request,
	})
	var createWorkOrderDTO workorder.CreateWorkOrder

	workOrderEntity, err := createWorkOrderDTO.ValidateAndGetEntity(ctx)
	if err != nil {
		log = log.WithField("error", err)
		log.Error()
		ctx.JSON(err.Code, gin.H{
			"error": err.Error.Error(),
		})
		return
	}

	log = log.WithField("workOrder", workOrderEntity)

	workOrder, err := c.CreateWorkOrderUC.Execute(*workOrderEntity)
	if err != nil {
		log = log.WithField("error", err)
		log.Error()
		ctx.JSON(err.Code, gin.H{
			"error": err.Error.Error(),
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
func (c *CustomerController) GetWorkOrders(ctx *gin.Context) {
	filters := dtos.WorkOrderFilters{
		ID: ctx.Param("id"),
	}

	workOrders := c.WorkOrderRepository.GetFiltered(filters)
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
func (c *CustomerController) GetCustomer(ctx *gin.Context) {
	log := c.Logger.WithFields(logrus.Fields{
		"file":    "customer_controller",
		"method":  "GetCustomer",
		"request": ctx.Request,
	})
	customer, err := c.CustomerRepository.FindByID(ctx.Param("id"))
	if err != nil {
		log = log.WithField("error", err)
		log.Error()
		ctx.JSON(err.Code, gin.H{
			"error": err.Error.Error(),
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
func (c *CustomerController) UpdateCustomer(ctx *gin.Context) {
	log := c.Logger.WithFields(logrus.Fields{
		"file":    "customer_controller",
		"method":  "UpdateCustomer",
		"request": ctx.Request,
	})
	var updateCustomerDTO customer.UpdateCustomer

	entity, err := updateCustomerDTO.ValidateAndGetEntity(ctx)
	if err != nil {
		log = log.WithField("error", err)
		log.Error()
		ctx.JSON(err.Code, gin.H{
			"error": err.Error.Error(),
		})
		return
	}

	entity.ID = ctx.Param("id")
	log = log.WithField("customerToUpdate", entity)

	updatedCustomer, err := c.UpdateCustomerUC.Execute(*entity)
	if err != nil {
		log = log.WithField("error", err)
		log.Error()
		ctx.JSON(err.Code, gin.H{
			"error": err.Error.Error(),
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
func (c *CustomerController) DeleteCustomer(ctx *gin.Context) {
	log := c.Logger.WithFields(logrus.Fields{
		"file":    "customer_controller",
		"method":  "DeleteCustomer",
		"request": ctx.Request,
	})

	if err := c.CustomerRepository.DeleteByID(ctx.Param("id")); err != nil {
		log = log.WithField("error", err)
		log.Error()
		ctx.JSON(err.Code, gin.H{
			"error": err.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
