package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"lucio.com/order-service/src/usecases/contracts"
)

type WorkOrderController struct {
	FinishWorkOrderUC contracts.FinishWorkOrderUC
}

func (w *WorkOrderController) GetWorkOrders(ctx *gin.Context) {

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
