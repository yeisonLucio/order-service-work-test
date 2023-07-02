package usecases

import (
	"lucio.com/order-service/src/dto"
	"lucio.com/order-service/src/repositories/contracts"
)

type UpdateWorkOrderUC struct {
	WorkOrderRepository contracts.WorkOrderRepository
}

func (u *UpdateWorkOrderUC) Execute(
	updateWorkOrder dto.UpdateWorkOrder,
) (*dto.UpdatedWorkOrder, error) {

	workOrder, err := u.WorkOrderRepository.FindByID(updateWorkOrder.ID)
	if err != nil {
		return nil, err
	}

	err = workOrder.SetPlannedDateBeginFromString(updateWorkOrder.PlannedDateBegin)
	if err != nil {
		return nil, err
	}

	err = workOrder.SetPlannedDateEndFromString(updateWorkOrder.PlannedDateEnd)
	if err != nil {
		return nil, err
	}

	workOrder.Title = updateWorkOrder.Title
	workOrder.Type = updateWorkOrder.Type

	if err := u.WorkOrderRepository.Save(workOrder); err != nil {
		return nil, err
	}

	return &dto.UpdatedWorkOrder{
		ID:               workOrder.ID.String(),
		Title:            workOrder.Title,
		PlannedDateBegin: workOrder.PlannedDateBegin.String(),
		PlannedDateEnd:   workOrder.PlannedDateEnd.String(),
		Type:             workOrder.Type,
		Status:           workOrder.Status,
	}, nil

}
