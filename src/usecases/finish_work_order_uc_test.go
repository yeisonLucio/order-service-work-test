package usecases

import (
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"lucio.com/order-service/src/mocks"
	"lucio.com/order-service/src/models"
)

func TestFinishWorkOrderUC_Execute(t *testing.T) {
	type fields struct {
		WorkOrderRepository *mocks.WorkOrderRepository
		CustomerRepository  *mocks.CustomerRepository
		EventRepository     *mocks.EventRepository
		Time                *mocks.Timer
	}
	type args struct {
		ID string
	}

	uuid := uuid.New()

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
			},
			args: args{
				ID: "uuid",
			},
			mocker: func(a args, f fields) {
				f.WorkOrderRepository.On("FindByID", a.ID).Return(nil, errors.New("error"))
			},
			wantErr: true,
		},
		{
			name: "should return an error when customer cannot be saved",
			fields: fields{
				WorkOrderRepository: &mocks.WorkOrderRepository{},
				CustomerRepository:  &mocks.CustomerRepository{},
				Time:                &mocks.Timer{},
			},
			args: args{
				ID: "uuid",
			},
			mocker: func(a args, f fields) {
				now := time.Now()
				f.Time.On("Now").Return(&now).Once()

				workOrder := models.WorkOrder{
					CustomerID: uuid,
				}
				f.WorkOrderRepository.On("FindByID", a.ID).Return(&workOrder, nil).Once()

				f.WorkOrderRepository.On(
					"IsTheFirstOrder",
					mock.Anything,
					mock.Anything,
				).Return(true).Once()

				customer := models.Customer{
					ID:        workOrder.CustomerID,
					IsActive:  true,
					StartDate: &now,
				}
				f.CustomerRepository.On("Save", &customer).Return(errors.New("error")).Once()
			},
			wantErr: true,
		},
		{
			name: "should return an error when work order cannot be saved",
			fields: fields{
				WorkOrderRepository: &mocks.WorkOrderRepository{},
				CustomerRepository:  &mocks.CustomerRepository{},
				Time:                &mocks.Timer{},
			},
			args: args{
				ID: "uuid",
			},
			mocker: func(a args, f fields) {
				now := time.Now()
				f.Time.On("Now").Return(&now).Once()

				workOrder := models.WorkOrder{
					CustomerID: uuid,
					Status:     models.StatusDone,
				}
				f.WorkOrderRepository.On("FindByID", a.ID).Return(&workOrder, nil).Once()

				f.WorkOrderRepository.On(
					"IsTheFirstOrder",
					mock.Anything,
					mock.Anything,
				).Return(false).Once()

				f.WorkOrderRepository.On(
					"Save",
					&workOrder,
				).Return(errors.New("error")).Once()
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
			},
			args: args{
				ID: "uuid",
			},
			mocker: func(a args, f fields) {
				now := time.Now()
				f.Time.On("Now").Return(&now).Once()

				workOrder := models.WorkOrder{
					CustomerID: uuid,
					Status:     models.StatusDone,
				}
				f.WorkOrderRepository.On("FindByID", a.ID).Return(&workOrder, nil).Once()

				f.WorkOrderRepository.On(
					"IsTheFirstOrder",
					mock.Anything,
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
			}
			if err := f.Execute(tt.args.ID); (err != nil) != tt.wantErr {
				t.Errorf("FinishWorkOrderUC.Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
