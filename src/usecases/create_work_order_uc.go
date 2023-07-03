package usecases

import (
	"errors"

	"lucio.com/order-service/src/dto"
	"lucio.com/order-service/src/helpers"
	"lucio.com/order-service/src/models"
	"lucio.com/order-service/src/repositories/contracts"
)

type CreateWorkOrderUC struct {
	WorkOrderRepository contracts.WorkOrderRepository
	UUID                helpers.UUIDGenerator
	CustomerRepository  contracts.CustomerRepository
	Time                helpers.Timer
}

func (c *CreateWorkOrderUC) Execute(
	createWorkOrderDTO dto.CreateWorkOrderDTO,
) (*dto.CreatedWorkOrderDTO, error) {
	customer, err := c.CustomerRepository.FindByID(createWorkOrderDTO.CustomerID)
	if err != nil {
		return nil, err
	}

	workOrder := models.WorkOrder{
		ID:         c.UUID.Generate(),
		CustomerID: customer.ID,
		Title:      createWorkOrderDTO.Title,
		Status:     models.StatusNew,
		Type:       createWorkOrderDTO.Type,
	}

	err = workOrder.SetPlannedDateBeginFromString(createWorkOrderDTO.PlannedDateBegin)
	if err != nil {
		return nil, err
	}

	err = workOrder.SetPlannedDateEndFromString(createWorkOrderDTO.PlannedDateEnd)
	if err != nil {
		return nil, err
	}

	if err := workOrder.Validate(); err != nil {
		return nil, err
	}

	if workOrder.Type == models.InactiveCustomerType {
		if !customer.IsActive {
			return nil, errors.New("el cliente ya se encuentra inactivo")
		}

		customer.IsActive = false
		customer.EndDate = c.Time.Now()

		if c.CustomerRepository.Save(customer) != nil {
			return nil, errors.New("el cliente no pudo ser actualizado")
		}
	}

	if err := c.WorkOrderRepository.Create(workOrder); err != nil {
		return nil, err
	}

	return &dto.CreatedWorkOrderDTO{
		ID:               workOrder.ID.String(),
		Title:            workOrder.Title,
		CustomerID:       workOrder.CustomerID.String(),
		PlannedDateBegin: workOrder.PlannedDateBegin.String(),
		PlannedDateEnd:   workOrder.PlannedDateEnd.String(),
		Type:             workOrder.Type,
		Status:           workOrder.Status,
	}, nil
}
