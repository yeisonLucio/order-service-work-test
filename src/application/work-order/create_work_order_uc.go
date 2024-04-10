package workorder

import (
	"errors"

	"github.com/sirupsen/logrus"
	"lucio.com/order-service/src/domain/common/dtos"
	"lucio.com/order-service/src/domain/common/helpers"
	consumerRepos "lucio.com/order-service/src/domain/customer/repositories"
	workOrderDtos "lucio.com/order-service/src/domain/workorder/dtos"
	"lucio.com/order-service/src/domain/workorder/entities"
	"lucio.com/order-service/src/domain/workorder/enums"
	"lucio.com/order-service/src/domain/workorder/repositories"
)

// CreateWorkOrderUC define las dependencias externas a utilizar  en este caso de uso
type CreateWorkOrderUC struct {
	WorkOrderRepository repositories.WorkOrderRepository
	CustomerRepository  consumerRepos.CustomerRepository
	Time                helpers.Timer
	Logger              *logrus.Logger
}

// Execute crea un nueva orden de servicio
func (c *CreateWorkOrderUC) Execute(
	createWorkOrder entities.WorkOrder,
) (*workOrderDtos.CreatedWorkOrderResponse, *dtos.CustomError) {
	log := c.Logger.WithFields(logrus.Fields{
		"file":   "create_work_order_uc",
		"method": "Execute",
	})
	customer, err := c.CustomerRepository.FindByID(createWorkOrder.CustomerID)
	if err != nil {
		log = log.WithField("error", err)
		log.Error()
		return nil, err
	}

	if createWorkOrder.PlannedDateBegin.After(*createWorkOrder.PlannedDateEnd) {
		log.Warning("la fecha de inicio no puede ser mayor que la fecha de fin")
		return nil, &dtos.CustomError{
			Code:  400,
			Error: errors.New("la fecha de inicio no puede ser mayor que la fecha de fin"),
		}
	}

	if createWorkOrder.Type == enums.InactiveCustomer {
		log.Warning("el cliente ya se encuentra inactivo")
		if !customer.IsActive {
			return nil, &dtos.CustomError{
				Code:  400,
				Error: errors.New("el cliente ya se encuentra inactivo"),
			}
		}

		customer.IsActive = false
		customer.EndDate = c.Time.Now()

		log = log.WithField("customer", customer)

		if err := c.CustomerRepository.Save(customer); err != nil {
			log = log.WithField("error", err)
			log.Error()
			return nil, err
		}
	}

	createWorkOrder.Status = enums.StatusNew

	log = log.WithField("workOrder", createWorkOrder)

	if err := c.WorkOrderRepository.Create(&createWorkOrder); err != nil {
		log = log.WithField("error", err)
		log.Error()
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
