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

func (w *WorkOrderController) GetWorkOrders(ctx *gin.Context) {
	filters := dto.WorkOrderFilters{
		PlannedDateBegin: ctx.Query("planned_date_begin"),
		PlannedDateEnd:   ctx.Query("planned_date_end"),
		Status:           ctx.Query("status"),
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": w.WorkOrderRepository.GetAll(filters),
	})
}

func (w *WorkOrderController) FinishWorkOrder(ctx *gin.Context) {
	if err := w.FinishWorkOrderUC.Execute(ctx.Param("id")); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
			"id":    "unexpected_error",
		})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (w *WorkOrderController) GetWorkOrder(ctx *gin.Context) {
	workOrder, err := w.WorkOrderRepository.FindByIdWithCustomer(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
			"id":    "record_not_found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": workOrder,
	})
}

func (w *WorkOrderController) UpdateWorkOrder(ctx *gin.Context) {
	var updateWorkOrder dto.UpdateWorkOrder

	if err := ctx.ShouldBindJSON(&updateWorkOrder); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"id":    "bad_request",
		})
		return
	}

	updateWorkOrder.ID = ctx.Param("id")

	updatedCustomer, err := w.UpdateWorkOrderUC.Execute(updateWorkOrder)
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

func (w *WorkOrderController) DeleteWorkOrder(ctx *gin.Context) {
	err := w.WorkOrderRepository.DeleteByID(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
			"id":    "record_not_found",
		})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
