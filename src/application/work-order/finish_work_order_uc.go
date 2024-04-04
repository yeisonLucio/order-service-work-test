package workorder

import (
	"lucio.com/order-service/src/domain/common/dtos"
	"lucio.com/order-service/src/domain/common/helpers"
	"lucio.com/order-service/src/domain/customer/entities"
	customerRepos "lucio.com/order-service/src/domain/customer/repositories"
	"lucio.com/order-service/src/domain/workorder/enums"
	"lucio.com/order-service/src/domain/workorder/repositories"
)

type FinishWorkOrderUC struct {
	WorkOrderRepository repositories.WorkOrderRepository
	CustomerRepository  customerRepos.CustomerRepository
	EventRepository     repositories.EventRepository
	Time                helpers.Timer
}

func (f *FinishWorkOrderUC) Execute(ID string) *dtos.CustomError {
	workOrder, err := f.WorkOrderRepository.FindByID(ID)
	if err != nil {
		return err
	}

	workOrder.Status = enums.StatusDone

	if f.WorkOrderRepository.IsTheFirstOrder(workOrder.CustomerID) {
		customer := entities.Customer{
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

	f.EventRepository.NotifyWorkOrderFinished(workOrder)

	return nil
}
