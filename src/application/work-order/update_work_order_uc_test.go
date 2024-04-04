package workorder

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	commonDtos "lucio.com/order-service/src/domain/common/dtos"
	"lucio.com/order-service/src/domain/workorder/dtos"
	"lucio.com/order-service/src/domain/workorder/entities"
	"lucio.com/order-service/src/domain/workorder/enums"
	"lucio.com/order-service/src/mocks"
)

func TestUpdateWorkOrderUC_Execute(t *testing.T) {
	type fields struct {
		WorkOrderRepository *mocks.WorkOrderRepository
	}
	type args struct {
		updateWorkOrder entities.WorkOrder
	}

	beginDate, _ := time.Parse(time.DateTime, "2023-01-01 01:00:00")
	endDate, _ := time.Parse(time.DateTime, "2023-01-01 02:00:00")

	tests := []struct {
		name    string
		fields  fields
		args    args
		mocker  func(a args, f fields)
		want    *dtos.UpdatedWorkOrderResponse
		wantErr bool
	}{
		{
			name: "should return an error when work order does not exist",
			fields: fields{
				WorkOrderRepository: &mocks.WorkOrderRepository{},
			},
			args: args{
				updateWorkOrder: entities.WorkOrder{},
			},
			mocker: func(a args, f fields) {
				f.WorkOrderRepository.On(
					"FindByID",
					a.updateWorkOrder.ID,
				).Return(nil, &commonDtos.CustomError{
					Error: errors.New("error"),
				})

			},
			wantErr: true,
		},
		{
			name: "should return an error when work order cannot be saved",
			fields: fields{
				WorkOrderRepository: &mocks.WorkOrderRepository{},
			},
			args: args{
				updateWorkOrder: entities.WorkOrder{
					PlannedDateBegin: &beginDate,
					PlannedDateEnd:   &endDate,
				},
			},
			mocker: func(a args, f fields) {
				f.WorkOrderRepository.On(
					"FindByID",
					a.updateWorkOrder.ID,
				).Return(&entities.WorkOrder{}, nil)

				f.WorkOrderRepository.On(
					"Save",
					mock.Anything,
				).Return(&commonDtos.CustomError{
					Error: errors.New("error"),
				})
			},
			wantErr: true,
		},
		{
			name: "should update a work order successfully",
			fields: fields{
				WorkOrderRepository: &mocks.WorkOrderRepository{},
			},
			args: args{
				updateWorkOrder: entities.WorkOrder{
					PlannedDateBegin: &beginDate,
					PlannedDateEnd:   &endDate,
				},
			},
			mocker: func(a args, f fields) {
				workOrder := entities.WorkOrder{
					ID:               "uuid",
					Title:            "test",
					PlannedDateBegin: nil,
					PlannedDateEnd:   nil,
					Type:             enums.ServiceOrder,
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
			want: &dtos.UpdatedWorkOrderResponse{
				ID:               "uuid",
				Title:            "test",
				PlannedDateBegin: beginDate.String(),
				PlannedDateEnd:   endDate.String(),
				Type:             string(enums.ServiceOrder),
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
