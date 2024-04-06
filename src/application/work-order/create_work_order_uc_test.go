package workorder

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/mock"
	commonDtos "lucio.com/order-service/src/domain/common/dtos"
	"lucio.com/order-service/src/domain/customer/entities"
	"lucio.com/order-service/src/domain/workorder/dtos"
	workOrderEntities "lucio.com/order-service/src/domain/workorder/entities"
	"lucio.com/order-service/src/domain/workorder/enums"
	"lucio.com/order-service/src/mocks"
)

func TestCreateWorkOrderUC_Execute(t *testing.T) {
	type fields struct {
		WorkOrderRepository *mocks.WorkOrderRepository
		CustomerRepository  *mocks.CustomerRepository
		Time                *mocks.Timer
		Logger              *logrus.Logger
	}
	type args struct {
		createWorkOrderDTO workOrderEntities.WorkOrder
	}

	beginDate := time.Now().Add(time.Hour * 24)
	endDate := time.Now().Add(time.Hour * 12)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mocker  func(a args, f fields)
		want    *dtos.CreatedWorkOrderResponse
		wantErr bool
	}{
		{
			name: "should return error when customer does not exist",
			fields: fields{
				CustomerRepository: &mocks.CustomerRepository{},
				Logger:             &logrus.Logger{},
			},
			args: args{
				createWorkOrderDTO: workOrderEntities.WorkOrder{
					CustomerID: "uuid",
				},
			},
			mocker: func(a args, f fields) {
				f.CustomerRepository.On("FindByID", "uuid").Return(nil, &commonDtos.CustomError{
					Code:  500,
					Error: errors.New("error"),
				})
			},
			wantErr: true,
		},
		{
			name: "should return error when begin date is greater than end date",
			fields: fields{
				CustomerRepository: &mocks.CustomerRepository{},
				Logger:             &logrus.Logger{},
			},
			args: args{
				createWorkOrderDTO: workOrderEntities.WorkOrder{
					CustomerID:       "uuid",
					PlannedDateBegin: &beginDate,
					PlannedDateEnd:   &endDate,
				},
			},
			mocker: func(a args, f fields) {
				f.CustomerRepository.On("FindByID", "uuid").
					Return(&entities.Customer{}, nil).
					Once()
			},
			wantErr: true,
		},
		{
			name: "should return error when customer is already inactive",
			fields: fields{
				CustomerRepository: &mocks.CustomerRepository{},
				Logger:             &logrus.Logger{},
			},
			args: args{
				createWorkOrderDTO: workOrderEntities.WorkOrder{
					Title:            "test",
					CustomerID:       "uuid",
					PlannedDateBegin: &beginDate,
					PlannedDateEnd:   &beginDate,
					Type:             enums.InactiveCustomer,
				},
			},
			mocker: func(a args, f fields) {
				f.CustomerRepository.On("FindByID", "uuid").
					Return(&entities.Customer{}, nil).
					Once()
			},
			wantErr: true,
		},
		{
			name: "should return error when the customer cannot be saved",
			fields: fields{
				CustomerRepository:  &mocks.CustomerRepository{},
				WorkOrderRepository: &mocks.WorkOrderRepository{},
				Time:                &mocks.Timer{},
				Logger:              &logrus.Logger{},
			},
			args: args{
				createWorkOrderDTO: workOrderEntities.WorkOrder{
					Title:            "test",
					CustomerID:       "uuid",
					PlannedDateBegin: &beginDate,
					PlannedDateEnd:   &beginDate,
					Type:             enums.InactiveCustomer,
				},
			},
			mocker: func(a args, f fields) {
				customer := &entities.Customer{
					IsActive: true,
				}
				f.CustomerRepository.On("FindByID", "uuid").
					Return(customer, nil).
					Once()
				now := time.Now()
				f.Time.On("Now").Return(&now).Once()
				f.CustomerRepository.On("Save", customer).Return(
					&commonDtos.CustomError{
						Code:  500,
						Error: errors.New("error"),
					},
				).Once()
			},
			wantErr: true,
		},
		{
			name: "should return error when the work order cannot be created",
			fields: fields{
				CustomerRepository:  &mocks.CustomerRepository{},
				WorkOrderRepository: &mocks.WorkOrderRepository{},
				Logger:              &logrus.Logger{},
			},
			args: args{
				createWorkOrderDTO: workOrderEntities.WorkOrder{
					Title:            "test",
					CustomerID:       "uuid",
					PlannedDateBegin: &beginDate,
					PlannedDateEnd:   &beginDate,
					Type:             enums.ServiceOrder,
				},
			},
			mocker: func(a args, f fields) {
				customer := &entities.Customer{}
				f.CustomerRepository.On("FindByID", "uuid").
					Return(customer, nil).
					Once()
				f.WorkOrderRepository.On("Create", mock.Anything).
					Return(
						&commonDtos.CustomError{
							Code:  500,
							Error: errors.New("error"),
						},
					).
					Once()
			},
			wantErr: true,
		},
		{
			name: "should create a work order successfully",
			fields: fields{
				CustomerRepository:  &mocks.CustomerRepository{},
				WorkOrderRepository: &mocks.WorkOrderRepository{},
				Logger:              &logrus.Logger{},
			},
			args: args{
				createWorkOrderDTO: workOrderEntities.WorkOrder{
					Title:            "test",
					CustomerID:       "uuid",
					PlannedDateBegin: &beginDate,
					PlannedDateEnd:   &beginDate,
					Type:             enums.ServiceOrder,
				},
			},
			mocker: func(a args, f fields) {
				customer := &entities.Customer{
					ID: "uuid",
				}
				f.CustomerRepository.On("FindByID", "uuid").
					Return(customer, nil).
					Once()

				workOrder := workOrderEntities.WorkOrder{
					CustomerID:       customer.ID,
					Title:            a.createWorkOrderDTO.Title,
					Status:           enums.StatusNew,
					Type:             a.createWorkOrderDTO.Type,
					PlannedDateBegin: a.createWorkOrderDTO.PlannedDateBegin,
					PlannedDateEnd:   a.createWorkOrderDTO.PlannedDateEnd,
				}

				f.WorkOrderRepository.On("Create", &workOrder).Return(nil).Once()
			},
			want: &dtos.CreatedWorkOrderResponse{
				Status:           string(enums.StatusNew),
				CustomerID:       "uuid",
				Title:            "test",
				PlannedDateBegin: beginDate.String(),
				PlannedDateEnd:   beginDate.String(),
				Type:             string(enums.ServiceOrder),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocker(tt.args, tt.fields)
			c := &CreateWorkOrderUC{
				WorkOrderRepository: tt.fields.WorkOrderRepository,
				CustomerRepository:  tt.fields.CustomerRepository,
				Time:                tt.fields.Time,
				Logger:              tt.fields.Logger,
			}
			got, err := c.Execute(tt.args.createWorkOrderDTO)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateWorkOrderUC.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateWorkOrderUC.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
