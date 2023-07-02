package usecases

import (
	"lucio.com/order-service/src/helpers"
	"lucio.com/order-service/src/models"
	"lucio.com/order-service/src/repositories/contracts"
)

type FinishWorkOrderUC struct {
	WorkOrderRepository contracts.WorkOrderRepository
	CustomerRepository  contracts.CustomerRepository
	EventRepository     contracts.EventRepository
	Time                helpers.Timer
}

func (f *FinishWorkOrderUC) Execute(ID string) error {
	workOrder, err := f.WorkOrderRepository.FindByID(ID)
	if err != nil {
		return err
	}

	workOrder.Status = models.StatusDone

	if f.WorkOrderRepository.IsTheFirstOrder(workOrder.ID.String(), workOrder.CustomerID.String()) {
		customer := models.Customer{
			ID:        workOrder.CustomerID,
			IsActive:  true,
			StartDate: f.Time.Now(),
		}

		if err := f.CustomerRepository.Save(&customer); err != nil {
			return err
		}
	}

	if err := f.WorkOrderRepository.Save(workOrder); err != nil {
		return err
	}

	f.EventRepository.NotifyWorkOrderFinished(*workOrder)

	return nil
}
