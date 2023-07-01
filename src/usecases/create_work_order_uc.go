package usecases

import (
	"errors"

	"lucio.com/order-service/src/dto"
	"lucio.com/order-service/src/helpers"
	"lucio.com/order-service/src/models"
	"lucio.com/order-service/src/repositories/contracts"
	"lucio.com/order-service/src/vo"
)

const limitDifference float64 = 2

type CreateWorkOrderUC struct {
	WorkOrderRepository contracts.WorkOrderRepository
	UUID                helpers.UUIDGenerator
	CustomerRepository  contracts.CustomerRepository
	Time                helpers.Timer
}

func (c *CreateWorkOrderUC) Execute(
	createWorkOrderDTO dto.CreateWorkOrderDTO,
) (*dto.CreatedWorkOrderDTO, error) {
	customer := c.CustomerRepository.FindByID(createWorkOrderDTO.CustomerID)
	if customer == nil {
		return nil, errors.New("el cliente ingresado no existe")
	}

	beginPlannedDate, err := c.Time.FromString(createWorkOrderDTO.PlannedDateBegin)
	if err != nil {
		return nil, errors.New("el formato de la fecha de inicio es incorrecto")
	}

	endPlannedDate, err := c.Time.FromString(createWorkOrderDTO.PlannedDateEnd)
	if err != nil {
		return nil, errors.New("el formato de la fecha de fin es incorrecto")
	}

	difference := endPlannedDate.Sub(beginPlannedDate)

	if difference.Hours() > limitDifference {
		return nil, errors.New("la diferencia de las fechas no puede ser mayor a dos horas")
	}

	var workOrderType vo.WorkOrderType

	if workOrderType.SetValue(createWorkOrderDTO.WorkOrderType) != nil {
		return nil, errors.New("el tipo de orden ingresada no esta permitida")
	}

	if workOrderType.GetValue() == vo.InactiveCustomerType && !customer.IsActive {
		return nil, errors.New("el cliente ya se encuentra inactivo")
	}

	workOrder := models.WorkOrder{
		ID:               c.UUID.Generate(),
		CustomerID:       customer.ID,
		Title:            createWorkOrderDTO.Title,
		PlannedDateBegin: beginPlannedDate,
		PlannedDateEnd:   endPlannedDate,
		Status:           vo.StatusNew,
		Type:             workOrderType,
	}

	if workOrderType.GetValue() == vo.InactiveCustomerType {
		customer.IsActive = false
		customer.EndDate = c.Time.Now()
		if c.CustomerRepository.Save(customer) != nil {
			return nil, errors.New("el cliente no pudo ser actualizado")
		}
	}

	c.WorkOrderRepository.Create(workOrder)

	return nil, nil
}
