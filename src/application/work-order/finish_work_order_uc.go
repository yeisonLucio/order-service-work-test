package workorder

import (
	"github.com/sirupsen/logrus"
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
	Logger              *logrus.Logger
}

func (f *FinishWorkOrderUC) Execute(ID string) *dtos.CustomError {
	log := f.Logger.WithFields(logrus.Fields{
		"file":   "finish_work_order_uc",
		"method": "Execute",
	})

	workOrder, err := f.WorkOrderRepository.FindByID(ID)
	if err != nil {
		log = log.WithField("error", err)
		log.Error()
		return err
	}

	workOrder.Status = enums.StatusDone

	if f.WorkOrderRepository.IsTheFirstOrder(workOrder.CustomerID) {
		customer := entities.Customer{
			ID:        workOrder.CustomerID,
			IsActive:  true,
			StartDate: f.Time.Now(),
		}

		log = log.WithField("customer", customer)

		if err := f.CustomerRepository.Save(&customer); err != nil {
			log = log.WithField("error", err)
			log.Error()
			return err
		}
	}

	log = log.WithField("workOrder", workOrder)

	if err := f.WorkOrderRepository.Save(workOrder); err != nil {
		log = log.WithField("error", err)
		log.Error()
		return err
	}

	f.EventRepository.NotifyWorkOrderFinished(workOrder)

	return nil
}
