package workorder

import (
	"errors"

	"lucio.com/order-service/src/domain/common/dtos"
	"lucio.com/order-service/src/domain/common/helpers"
	consumerRepos "lucio.com/order-service/src/domain/customer/repositories"
	workOrderDtos "lucio.com/order-service/src/domain/workorder/dtos"
	"lucio.com/order-service/src/domain/workorder/entities"
	"lucio.com/order-service/src/domain/workorder/enums"
	"lucio.com/order-service/src/domain/workorder/repositories"
)

type CreateWorkOrderUC struct {
	WorkOrderRepository repositories.WorkOrderRepository
	CustomerRepository  consumerRepos.CustomerRepository
	Time                helpers.Timer
}

func (c *CreateWorkOrderUC) Execute(
	createWorkOrder entities.WorkOrder,
) (*workOrderDtos.CreatedWorkOrderResponse, *dtos.CustomError) {
	customer, err := c.CustomerRepository.FindByID(createWorkOrder.CustomerID)
	if err != nil {
		return nil, err
	}

	if createWorkOrder.PlannedDateBegin.After(*createWorkOrder.PlannedDateEnd) {
		return nil, &dtos.CustomError{
			Code:  400,
			Error: errors.New("la fecha de inicio no puede ser mayor que la fecha de fin"),
		}
	}

	if createWorkOrder.Type == enums.InactiveCustomer {
		if !customer.IsActive {
			return nil, &dtos.CustomError{
				Code:  400,
				Error: errors.New("el cliente ya se encuentra inactivo"),
			}
		}

		customer.IsActive = false
		customer.EndDate = c.Time.Now()

		if err := c.CustomerRepository.Save(customer); err != nil {
			return nil, err
		}
	}

	createWorkOrder.Status = enums.StatusNew

	if err := c.WorkOrderRepository.Create(&createWorkOrder); err != nil {
		return nil, err
	}

	return &workOrderDtos.CreatedWorkOrderResponse{
		ID:               createWorkOrder.ID,
		Title:            createWorkOrder.Title,
		CustomerID:       createWorkOrder.CustomerID,
		PlannedDateBegin: createWorkOrder.PlannedDateBegin.String(),
		PlannedDateEnd:   createWorkOrder.PlannedDateEnd.String(),
		Type:             string(createWorkOrder.Type),
		Status:           string(createWorkOrder.Status),
	}, nil
}
