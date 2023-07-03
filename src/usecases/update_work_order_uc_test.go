package usecases

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"lucio.com/order-service/src/dto"
	"lucio.com/order-service/src/mocks"
	"lucio.com/order-service/src/models"
)

func TestUpdateWorkOrderUC_Execute(t *testing.T) {
	type fields struct {
		WorkOrderRepository *mocks.WorkOrderRepository
	}
	type args struct {
		updateWorkOrder dto.UpdateWorkOrder
	}

	uuid := uuid.New()
	beginDate, _ := time.Parse(time.DateTime, "2023-01-01 01:00:00")
	endDate, _ := time.Parse(time.DateTime, "2023-01-01 02:00:00")

	tests := []struct {
		name    string
		fields  fields
		args    args
		mocker  func(a args, f fields)
		want    *dto.UpdatedWorkOrder
		wantErr bool
	}{
		{
			name: "should return an error when work order does not exist",
			fields: fields{
				WorkOrderRepository: &mocks.WorkOrderRepository{},
			},
			args: args{
				updateWorkOrder: dto.UpdateWorkOrder{},
			},
			mocker: func(a args, f fields) {
				f.WorkOrderRepository.On(
					"FindByID",
					a.updateWorkOrder.ID,
				).Return(nil, errors.New("error"))

			},
			wantErr: true,
		},
		{
			name: "should return an error when begin date is not valid",
			fields: fields{
				WorkOrderRepository: &mocks.WorkOrderRepository{},
			},
			args: args{
				updateWorkOrder: dto.UpdateWorkOrder{},
			},
			mocker: func(a args, f fields) {
				f.WorkOrderRepository.On(
					"FindByID",
					a.updateWorkOrder.ID,
				).Return(&models.WorkOrder{}, nil)
			},
			wantErr: true,
		},
		{
			name: "should return an error when end date is not valid",
			fields: fields{
				WorkOrderRepository: &mocks.WorkOrderRepository{},
			},
			args: args{
				updateWorkOrder: dto.UpdateWorkOrder{
					PlannedDateBegin: "2023-01-01 01:00:00",
				},
			},
			mocker: func(a args, f fields) {
				f.WorkOrderRepository.On(
					"FindByID",
					a.updateWorkOrder.ID,
				).Return(&models.WorkOrder{}, nil)
			},
			wantErr: true,
		},
		{
			name: "should return an error when dates are not valid",
			fields: fields{
				WorkOrderRepository: &mocks.WorkOrderRepository{},
			},
			args: args{
				updateWorkOrder: dto.UpdateWorkOrder{
					PlannedDateBegin: "2023-01-01 01:00:00",
					PlannedDateEnd:   "2023-01-01 04:00:00",
				},
			},
			mocker: func(a args, f fields) {
				f.WorkOrderRepository.On(
					"FindByID",
					a.updateWorkOrder.ID,
				).Return(&models.WorkOrder{}, nil)
			},
			wantErr: true,
		},
		{
			name: "should return an error when work order cannot be saved",
			fields: fields{
				WorkOrderRepository: &mocks.WorkOrderRepository{},
			},
			args: args{
				updateWorkOrder: dto.UpdateWorkOrder{
					PlannedDateBegin: "2023-01-01 01:00:00",
					PlannedDateEnd:   "2023-01-01 02:00:00",
				},
			},
			mocker: func(a args, f fields) {
				f.WorkOrderRepository.On(
					"FindByID",
					a.updateWorkOrder.ID,
				).Return(&models.WorkOrder{}, nil)

				f.WorkOrderRepository.On(
					"Save",
					mock.Anything,
				).Return(errors.New("error"))
			},
			wantErr: true,
		},
		{
			name: "should update a work order successfully",
			fields: fields{
				WorkOrderRepository: &mocks.WorkOrderRepository{},
			},
			args: args{
				updateWorkOrder: dto.UpdateWorkOrder{
					PlannedDateBegin: "2023-01-01 01:00:00",
					PlannedDateEnd:   "2023-01-01 02:00:00",
				},
			},
			mocker: func(a args, f fields) {
				workOrder := models.WorkOrder{
					ID:               uuid,
					Title:            "test",
					PlannedDateBegin: nil,
					PlannedDateEnd:   nil,
					Type:             models.ServiceOrderType,
				}
				f.WorkOrderRepository.On(
					"FindByID",
					a.updateWorkOrder.ID,
				).Return(&workOrder, nil)

				workOrder.PlannedDateBegin = &beginDate
				workOrder.PlannedDateEnd = &endDate

				f.WorkOrderRepository.On(
					"Save",
					&workOrder,
				).Return(nil)
			},
			want: &dto.UpdatedWorkOrder{
				ID:               uuid.String(),
				Title:            "test",
				PlannedDateBegin: beginDate.String(),
				PlannedDateEnd:   endDate.String(),
				Type:             models.ServiceOrderType,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocker(tt.args, tt.fields)
			u := &UpdateWorkOrderUC{
				WorkOrderRepository: tt.fields.WorkOrderRepository,
			}
			got, err := u.Execute(tt.args.updateWorkOrder)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateWorkOrderUC.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateWorkOrderUC.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
