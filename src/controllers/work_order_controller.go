package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"lucio.com/order-service/src/dto"
	rcontracts "lucio.com/order-service/src/repositories/contracts"
	"lucio.com/order-service/src/usecases/contracts"
)

type WorkOrderController struct {
	FinishWorkOrderUC   contracts.FinishWorkOrderUC
	WorkOrderRepository rcontracts.WorkOrderRepository
	UpdateWorkOrderUC   contracts.UpdateWorkOrderUC
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
	filters := dto.WorkOrderFilters{
		PlannedDateBegin: ctx.Query("planned_date_begin"),
		PlannedDateEnd:   ctx.Query("planned_date_end"),
		Status:           ctx.Query("status"),
	}

	ctx.JSON(http.StatusOK, w.WorkOrderRepository.GetAll(filters))
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
	if err := w.FinishWorkOrderUC.Execute(ctx.Param("id")); err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			ID:      "unexpected_error",
			Message: err.Error(),
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
	workOrder, err := w.WorkOrderRepository.FindByIdWithCustomer(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.ErrorResponse{
			ID:      "record_not_found",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, workOrder)
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
	var updateWorkOrder dto.UpdateWorkOrder

	if err := ctx.ShouldBindJSON(&updateWorkOrder); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			ID:      "bad_request",
			Message: err.Error(),
		})
		return
	}

	updateWorkOrder.ID = ctx.Param("id")

	updatedCustomer, err := w.UpdateWorkOrderUC.Execute(updateWorkOrder)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			ID:      "unexpected_error",
			Message: err.Error(),
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
	err := w.WorkOrderRepository.DeleteByID(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.ErrorResponse{
			ID:      "record_not_found",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
