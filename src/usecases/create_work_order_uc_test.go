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

func TestCreateWorkOrderUC_Execute(t *testing.T) {
	type fields struct {
		WorkOrderRepository *mocks.WorkOrderRepository
		UUID                *mocks.UUIDGenerator
		CustomerRepository  *mocks.CustomerRepository
		Time                *mocks.Timer
	}
	uuid := uuid.New()
	type args struct {
		createWorkOrderDTO dto.CreateWorkOrderDTO
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mocker  func(a args, f fields)
		want    *dto.CreatedWorkOrderDTO
		wantErr bool
	}{
		{
			name: "should return error when customer does not exist",
			fields: fields{
				CustomerRepository: &mocks.CustomerRepository{},
			},
			args: args{
				createWorkOrderDTO: dto.CreateWorkOrderDTO{
					CustomerID: "uuid",
				},
			},
			mocker: func(a args, f fields) {
				f.CustomerRepository.On("FindByID", "uuid").Return(nil, errors.New("error"))
			},
			wantErr: true,
		},
		{
			name: "should return error when begin date is not valid",
			fields: fields{
				CustomerRepository: &mocks.CustomerRepository{},
				UUID:               &mocks.UUIDGenerator{},
			},
			args: args{
				createWorkOrderDTO: dto.CreateWorkOrderDTO{
					CustomerID: "uuid",
				},
			},
			mocker: func(a args, f fields) {
				f.CustomerRepository.On("FindByID", "uuid").
					Return(&models.Customer{}, nil).
					Once()
				f.UUID.On("Generate").Return(uuid).Once()
			},
			wantErr: true,
		},
		{
			name: "should return error when end date is not valid",
			fields: fields{
				CustomerRepository: &mocks.CustomerRepository{},
				UUID:               &mocks.UUIDGenerator{},
			},
			args: args{
				createWorkOrderDTO: dto.CreateWorkOrderDTO{
					CustomerID:       "uuid",
					PlannedDateBegin: "2023-01-01 00:00:00",
				},
			},
			mocker: func(a args, f fields) {
				f.CustomerRepository.On("FindByID", "uuid").
					Return(&models.Customer{}, nil).
					Once()
				f.UUID.On("Generate").Return(uuid).Once()
			},
			wantErr: true,
		},
		{
			name: "should return error when type is not valid",
			fields: fields{
				CustomerRepository: &mocks.CustomerRepository{},
				UUID:               &mocks.UUIDGenerator{},
			},
			args: args{
				createWorkOrderDTO: dto.CreateWorkOrderDTO{
					Title:            "test",
					CustomerID:       "uuid",
					PlannedDateBegin: "2023-01-01 00:00:00",
					PlannedDateEnd:   "2023-01-01 00:00:00",
					Type:             "fake",
				},
			},
			mocker: func(a args, f fields) {
				f.CustomerRepository.On("FindByID", "uuid").
					Return(&models.Customer{}, nil).
					Once()
				f.UUID.On("Generate").Return(uuid).Once()
			},
			wantErr: true,
		},
		{
			name: "should return error when customer is already inactive",
			fields: fields{
				CustomerRepository: &mocks.CustomerRepository{},
				UUID:               &mocks.UUIDGenerator{},
			},
			args: args{
				createWorkOrderDTO: dto.CreateWorkOrderDTO{
					Title:            "test",
					CustomerID:       "uuid",
					PlannedDateBegin: "2023-01-01 00:00:00",
					PlannedDateEnd:   "2023-01-01 00:00:00",
					Type:             models.InactiveCustomerType,
				},
			},
			mocker: func(a args, f fields) {
				f.CustomerRepository.On("FindByID", "uuid").
					Return(&models.Customer{}, nil).
					Once()
				f.UUID.On("Generate").Return(uuid).Once()
			},
			wantErr: true,
		},
		{
			name: "should return error when the customer cannot be saved",
			fields: fields{
				CustomerRepository:  &mocks.CustomerRepository{},
				UUID:                &mocks.UUIDGenerator{},
				WorkOrderRepository: &mocks.WorkOrderRepository{},
				Time:                &mocks.Timer{},
			},
			args: args{
				createWorkOrderDTO: dto.CreateWorkOrderDTO{
					Title:            "test",
					CustomerID:       "uuid",
					PlannedDateBegin: "2023-01-01 00:00:00",
					PlannedDateEnd:   "2023-01-01 00:00:00",
					Type:             models.InactiveCustomerType,
				},
			},
			mocker: func(a args, f fields) {
				customer := &models.Customer{
					IsActive: true,
				}
				f.CustomerRepository.On("FindByID", "uuid").
					Return(customer, nil).
					Once()
				f.UUID.On("Generate").Return(uuid).Once()
				now := time.Now()
				f.Time.On("Now").Return(&now).Once()
				f.CustomerRepository.On("Save", customer).Return(errors.New("error")).Once()
			},
			wantErr: true,
		},
		{
			name: "should return error when the work order cannot be created",
			fields: fields{
				CustomerRepository:  &mocks.CustomerRepository{},
				UUID:                &mocks.UUIDGenerator{},
				WorkOrderRepository: &mocks.WorkOrderRepository{},
			},
			args: args{
				createWorkOrderDTO: dto.CreateWorkOrderDTO{
					Title:            "test",
					CustomerID:       "uuid",
					PlannedDateBegin: "2023-01-01 00:00:00",
					PlannedDateEnd:   "2023-01-01 00:00:00",
					Type:             models.ServiceOrderType,
				},
			},
			mocker: func(a args, f fields) {
				customer := &models.Customer{}
				f.CustomerRepository.On("FindByID", "uuid").
					Return(customer, nil).
					Once()
				f.UUID.On("Generate").Return(uuid).Once()
				f.WorkOrderRepository.On("Create", mock.Anything).
					Return(errors.New("error")).
					Once()
			},
			wantErr: true,
		},
		{
			name: "should create a work order successfully",
			fields: fields{
				CustomerRepository:  &mocks.CustomerRepository{},
				UUID:                &mocks.UUIDGenerator{},
				WorkOrderRepository: &mocks.WorkOrderRepository{},
			},
			args: args{
				createWorkOrderDTO: dto.CreateWorkOrderDTO{
					Title:            "test",
					CustomerID:       "uuid",
					PlannedDateBegin: "2023-01-01 10:00:00",
					PlannedDateEnd:   "2023-01-01 11:00:00",
					Type:             models.ServiceOrderType,
				},
			},
			mocker: func(a args, f fields) {
				customer := &models.Customer{
					ID: uuid,
				}
				f.CustomerRepository.On("FindByID", "uuid").
					Return(customer, nil).
					Once()
				f.UUID.On("Generate").Return(uuid).Once()

				workOrder := models.WorkOrder{
					ID:         uuid,
					CustomerID: customer.ID,
					Title:      a.createWorkOrderDTO.Title,
					Status:     models.StatusNew,
					Type:       a.createWorkOrderDTO.Type,
				}
				workOrder.SetPlannedDateBegin(a.createWorkOrderDTO.PlannedDateBegin)
				workOrder.SetPlannedDateEnd(a.createWorkOrderDTO.PlannedDateEnd)

				f.WorkOrderRepository.On("Create", workOrder).Return(nil).Once()
			},
			want: &dto.CreatedWorkOrderDTO{
				ID:               uuid.String(),
				Status:           models.StatusNew,
				CustomerID:       uuid.String(),
				Title:            "test",
				PlannedDateBegin: "2023-01-01 10:00:00 +0000 UTC",
				PlannedDateEnd:   "2023-01-01 11:00:00 +0000 UTC",
				Type:             models.ServiceOrderType,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocker(tt.args, tt.fields)
			c := &CreateWorkOrderUC{
				WorkOrderRepository: tt.fields.WorkOrderRepository,
				UUID:                tt.fields.UUID,
				CustomerRepository:  tt.fields.CustomerRepository,
				Time:                tt.fields.Time,
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
