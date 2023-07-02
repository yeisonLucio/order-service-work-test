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
