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

	err = workOrder.SetPlannedDateBegin(updateWorkOrder.PlannedDateBegin)
	if err != nil {
		return nil, err
	}

	err = workOrder.SetPlannedDateEnd(updateWorkOrder.PlannedDateEnd)
	if err != nil {
		return nil, err
	}

	if updateWorkOrder.Title != "" {
		workOrder.Title = updateWorkOrder.Title
	}

	if err := workOrder.Validate(); err != nil {
		return nil, err
	}

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
