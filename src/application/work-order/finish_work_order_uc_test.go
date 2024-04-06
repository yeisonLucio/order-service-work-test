package workorder

import (
	"errors"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/mock"
	"lucio.com/order-service/src/domain/common/dtos"
	customerEntities "lucio.com/order-service/src/domain/customer/entities"
	"lucio.com/order-service/src/domain/workorder/entities"
	"lucio.com/order-service/src/domain/workorder/enums"
	"lucio.com/order-service/src/mocks"
)

func TestFinishWorkOrderUC_Execute(t *testing.T) {
	type fields struct {
		WorkOrderRepository *mocks.WorkOrderRepository
		CustomerRepository  *mocks.CustomerRepository
		EventRepository     *mocks.EventRepository
		Time                *mocks.Timer
		Logger              *logrus.Logger
	}
	type args struct {
		ID string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		mocker  func(a args, f fields)
		wantErr bool
	}{
		{
			name: "should return an error when work order id does not exist",
			fields: fields{
				WorkOrderRepository: &mocks.WorkOrderRepository{},
				Logger:              &logrus.Logger{},
			},
			args: args{
				ID: "uuid",
			},
			mocker: func(a args, f fields) {
				f.WorkOrderRepository.On("FindByID", a.ID).Return(nil, &dtos.CustomError{
					Error: errors.New("error"),
				})
			},
			wantErr: true,
		},
		{
			name: "should return an error when customer cannot be saved",
			fields: fields{
				WorkOrderRepository: &mocks.WorkOrderRepository{},
				CustomerRepository:  &mocks.CustomerRepository{},
				Time:                &mocks.Timer{},
				Logger:              &logrus.Logger{},
			},
			args: args{
				ID: "uuid",
			},
			mocker: func(a args, f fields) {
				now := time.Now()
				f.Time.On("Now").Return(&now).Once()

				workOrder := entities.WorkOrder{
					CustomerID: "uuid",
				}
				f.WorkOrderRepository.On("FindByID", a.ID).Return(&workOrder, nil).Once()

				f.WorkOrderRepository.On(
					"IsTheFirstOrder",
					mock.Anything,
				).Return(true).Once()

				customer := customerEntities.Customer{
					ID:        workOrder.CustomerID,
					IsActive:  true,
					StartDate: &now,
				}
				f.CustomerRepository.On("Save", &customer).Return(&dtos.CustomError{
					Error: errors.New("error"),
				}).Once()
			},
			wantErr: true,
		},
		{
			name: "should return an error when work order cannot be saved",
			fields: fields{
				WorkOrderRepository: &mocks.WorkOrderRepository{},
				CustomerRepository:  &mocks.CustomerRepository{},
				Time:                &mocks.Timer{},
				Logger:              &logrus.Logger{},
			},
			args: args{
				ID: "uuid",
			},
			mocker: func(a args, f fields) {
				now := time.Now()
				f.Time.On("Now").Return(&now).Once()

				workOrder := entities.WorkOrder{
					CustomerID: "uuid",
					Status:     enums.StatusDone,
				}
				f.WorkOrderRepository.On("FindByID", a.ID).Return(&workOrder, nil).Once()

				f.WorkOrderRepository.On(
					"IsTheFirstOrder",
					mock.Anything,
				).Return(false).Once()

				f.WorkOrderRepository.On(
					"Save",
					&workOrder,
				).Return(&dtos.CustomError{
					Error: errors.New("error"),
				}).Once()
			},
			wantErr: true,
		},
		{
			name: "should return an error when work order cannot be saved",
			fields: fields{
				WorkOrderRepository: &mocks.WorkOrderRepository{},
				CustomerRepository:  &mocks.CustomerRepository{},
				Time:                &mocks.Timer{},
				EventRepository:     &mocks.EventRepository{},
				Logger:              &logrus.Logger{},
			},
			args: args{
				ID: "uuid",
			},
			mocker: func(a args, f fields) {
				now := time.Now()
				f.Time.On("Now").Return(&now).Once()

				workOrder := entities.WorkOrder{
					CustomerID: "uuid",
					Status:     enums.StatusDone,
				}
				f.WorkOrderRepository.On("FindByID", a.ID).Return(&workOrder, nil).Once()

				f.WorkOrderRepository.On(
					"IsTheFirstOrder",
					mock.Anything,
				).Return(false).Once()

				f.WorkOrderRepository.On(
					"Save",
					&workOrder,
				).Return(nil).Once()

				f.EventRepository.On(
					"NotifyWorkOrderFinished",
					&workOrder,
				).Return(nil).Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocker(tt.args, tt.fields)
			f := &FinishWorkOrderUC{
				WorkOrderRepository: tt.fields.WorkOrderRepository,
				CustomerRepository:  tt.fields.CustomerRepository,
				EventRepository:     tt.fields.EventRepository,
				Time:                tt.fields.Time,
				Logger:              tt.fields.Logger,
			}
			if err := f.Execute(tt.args.ID); (err != nil) != tt.wantErr {
				t.Errorf("FinishWorkOrderUC.Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
