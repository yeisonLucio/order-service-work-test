package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"lucio.com/order-service/src/domain/workorder/dtos"
	"lucio.com/order-service/src/domain/workorder/repositories"
	usecases "lucio.com/order-service/src/domain/workorder/usecases"
	"lucio.com/order-service/src/infra/requests/workorder"
)

type WorkOrderController struct {
	FinishWorkOrderUC   usecases.FinishWorkOrderUC
	WorkOrderRepository repositories.WorkOrderRepository
	UpdateWorkOrderUC   usecases.UpdateWorkOrderUC
	Logger              *logrus.Logger
}

// @Summary Servicio para obtener las ordenes de servicio tendiendo en cuenta los filtros
// @Description Este servicio permite filtrar las ordenes de servicio
// @Tags WorkOrders
// @Accept json
// @Produce json
// @param planned_date_begin query string true "fecha de planeación inicial"
// @param planned_date_end query string true "fecha de planeación final"
// @param status query string true "Estado de la orden de servicio"
// @Success 200 {array} dto.WorkOrderDTO
// @Router /work-orders [get]
func (w *WorkOrderController) GetWorkOrders(ctx *gin.Context) {
	filters := dtos.WorkOrderFilters{
		PlannedDateBegin: ctx.Query("planned_date_begin"),
		PlannedDateEnd:   ctx.Query("planned_date_end"),
		Status:           ctx.Query("status"),
	}

	ctx.JSON(http.StatusOK, w.WorkOrderRepository.GetFiltered(filters))
}

// @Summary Servicio para finalizar una orden de servicio
// @Description Este servicio permite finalizar una orden de servicio
// @Tags WorkOrders
// @Accept json
// @Produce json
// @param id path string true "id de la orden de servicio"
// @Success 204 "No content"
// @Failure 500 {object} dto.ErrorResponse
// @Router /work-orders/{id}/finish [patch]
func (w *WorkOrderController) FinishWorkOrder(ctx *gin.Context) {
	log := w.Logger.WithFields(logrus.Fields{
		"file":    "work_order_controller",
		"method":  "FinishWorkOrder",
		"request": ctx.Request,
	})
	if err := w.FinishWorkOrderUC.Execute(ctx.Param("id")); err != nil {
		log = log.WithField("error", err)
		log.Error()
		ctx.JSON(err.Code, gin.H{
			"error": err.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

// @Summary Servicio para obtener una orden de servicio
// @Description Este servicio permite obtener una orden de servicio
// @Tags WorkOrders
// @Accept json
// @Produce json
// @param id path string true "id de la orden de servicio"
// @Success 200 {object} dto.WorkOrderDTO
// @Failure 404 {object} dto.ErrorResponse
// @Router /work-orders/{id} [get]
func (w *WorkOrderController) GetWorkOrder(ctx *gin.Context) {
	filters := dtos.WorkOrderFilters{
		ID: ctx.Param("id"),
	}

	workOrders := w.WorkOrderRepository.GetFiltered(filters)

	if len(workOrders) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "orden de servicio no encontrada",
		})
		return
	}

	ctx.JSON(http.StatusOK, workOrders[0])
}

// @Summary Servicio para actualizar una orden de servicio
// @Description Este servicio permite actualizar una orden de servicio
// @Tags WorkOrders
// @Accept json
// @Produce json
// @param id path string true "id de la orden de servicio"
// @Param body body dto.UpdateWorkOrder true "Body data"
// @Success 200 {object} dto.UpdatedWorkOrder
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /work-orders/{id} [put]
func (w *WorkOrderController) UpdateWorkOrder(ctx *gin.Context) {
	log := w.Logger.WithFields(logrus.Fields{
		"file":    "work_order_controller",
		"method":  "UpdateWorkOrder",
		"request": ctx.Request,
	})

	var updateWorkOrder workorder.UpdateWorkOrder

	workOrderEntity, err := updateWorkOrder.ValidateAndGetEntity(ctx)
	if err != nil {
		log = log.WithField("error", err)
		log.Error()
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error.Error(),
		})
		return
	}

	log = log.WithField("workOrder", workOrderEntity)

	updatedCustomer, err := w.UpdateWorkOrderUC.Execute(*workOrderEntity)
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

// @Summary Servicio para eliminar una orden de servicio
// @Description Este servicio permite eliminar una orden de servicio
// @Tags WorkOrders
// @Accept json
// @Produce json
// @param id path string true "id de la orden de servicio"
// @Success 204 "No content"
// @Failure 404 {object} dto.ErrorResponse
// @Router /work-orders/{id} [delete]
func (w *WorkOrderController) DeleteWorkOrder(ctx *gin.Context) {
	log := w.Logger.WithFields(logrus.Fields{
		"file":    "customer_controller",
		"method":  "DeleteWorkOrder",
		"request": ctx.Request,
	})
	err := w.WorkOrderRepository.DeleteByID(ctx.Param("id"))
	if err != nil {
		log = log.WithField("error", err)
		log.Error()
		ctx.JSON(err.Code, gin.H{
			"error": err.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
