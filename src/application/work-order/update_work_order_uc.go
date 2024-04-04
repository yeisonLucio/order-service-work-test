package workorder

import (
	"lucio.com/order-service/src/domain/common/dtos"
	workOrderDtos "lucio.com/order-service/src/domain/workorder/dtos"
	"lucio.com/order-service/src/domain/workorder/entities"
	"lucio.com/order-service/src/domain/workorder/repositories"
)

type UpdateWorkOrderUC struct {
	WorkOrderRepository repositories.WorkOrderRepository
}

func (u *UpdateWorkOrderUC) Execute(
	updateWorkOrder entities.WorkOrder,
) (*workOrderDtos.UpdatedWorkOrderResponse, *dtos.CustomError) {
	workOrder, err := u.WorkOrderRepository.FindByID(updateWorkOrder.ID)
	if err != nil {
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

	if err := u.WorkOrderRepository.Save(workOrder); err != nil {
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
