package workorder

import (
	"github.com/sirupsen/logrus"
	"lucio.com/order-service/src/domain/common/dtos"
	workOrderDtos "lucio.com/order-service/src/domain/workorder/dtos"
	"lucio.com/order-service/src/domain/workorder/entities"
	"lucio.com/order-service/src/domain/workorder/repositories"
)

// UpdateWorkOrderUC define las dependencias externas a utilizar
type UpdateWorkOrderUC struct {
	WorkOrderRepository repositories.WorkOrderRepository
	Logger              *logrus.Logger
}

// Execute permite actualizar una orden de servicio
func (u *UpdateWorkOrderUC) Execute(
	updateWorkOrder entities.WorkOrder,
) (*workOrderDtos.UpdatedWorkOrderResponse, *dtos.CustomError) {
	log := u.Logger.WithFields(logrus.Fields{
		"file":   "update_work_order_uc",
		"method": "Execute",
	})
	workOrder, err := u.WorkOrderRepository.FindByID(updateWorkOrder.ID)
	if err != nil {
		log = log.WithField("error", err)
		log.Error()
		return nil, err
	}

	if updateWorkOrder.Title != "" {
		workOrder.Title = updateWorkOrder.Title
	}

	if updateWorkOrder.PlannedDateBegin != nil {
		workOrder.PlannedDateBegin = updateWorkOrder.PlannedDateBegin
	}

	if updateWorkOrder.PlannedDateEnd != nil {
		workOrder.PlannedDateEnd = updateWorkOrder.PlannedDateEnd
	}

	log = log.WithField("workOrder", workOrder)

	if err := u.WorkOrderRepository.Save(workOrder); err != nil {
		log = log.WithField("error", err)
		log.Error()
		return nil, err
	}

	return &workOrderDtos.UpdatedWorkOrderResponse{
		ID:               workOrder.ID,
		Title:            workOrder.Title,
		PlannedDateBegin: workOrder.PlannedDateBegin.String(),
		PlannedDateEnd:   workOrder.PlannedDateEnd.String(),
		Type:             string(workOrder.Type),
		Status:           string(workOrder.Status),
	}, nil
}
