package usecases

import (
	"errors"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"lucio.com/order-service/src/dto"
	"lucio.com/order-service/src/mocks"
	"lucio.com/order-service/src/models"
)

func TestCreateCustomerUC_Execute(t *testing.T) {
	type fields struct {
		CustomerRepository *mocks.CustomerRepository
		UUID               *mocks.UUIDGenerator
	}
	type args struct {
		createCustomerDTO dto.CreateCustomerDTO
	}
	uuid := uuid.New()
	tests := []struct {
		name    string
		fields  fields
		args    args
		mocker  func(a args, f fields)
		want    *dto.CreatedCustomerDTO
		wantErr bool
	}{
		{
			name: "should create customer successfully",
			fields: fields{
				CustomerRepository: &mocks.CustomerRepository{},
				UUID:               &mocks.UUIDGenerator{},
			},
			args: args{
				createCustomerDTO: dto.CreateCustomerDTO{
					FirstName: "John",
					LastName:  "Doe",
					Address:   "123",
				},
			},
			mocker: func(a args, f fields) {
				f.UUID.On("Generate").Return(uuid).Once()
				customer := models.Customer{
					ID:        uuid,
					FirstName: "John",
					LastName:  "Doe",
					Address:   "123",
				}
				f.CustomerRepository.On("Create", customer).Return(nil).Once()
			},
			want: &dto.CreatedCustomerDTO{
				ID:        uuid.String(),
				IsActive:  false,
				FirstName: "John",
				LastName:  "Doe",
				Address:   "123",
			},
		},
		{
			name: "should return an error when customer is no valid",
			fields: fields{
				UUID: &mocks.UUIDGenerator{},
			},
			args: args{
				createCustomerDTO: dto.CreateCustomerDTO{},
			},
			mocker: func(a args, f fields) {
				f.UUID.On("Generate").Return(uuid).Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should return an error when customer cannot be created",
			fields: fields{
				CustomerRepository: &mocks.CustomerRepository{},
				UUID:               &mocks.UUIDGenerator{},
			},
			args: args{
				createCustomerDTO: dto.CreateCustomerDTO{
					FirstName: "John",
					LastName:  "Doe",
					Address:   "123",
				},
			},
			mocker: func(a args, f fields) {
				f.UUID.On("Generate").Return(uuid).Once()
				f.CustomerRepository.On("Create", models.Customer{
					ID:        uuid,
					FirstName: "John",
					LastName:  "Doe",
					Address:   "123",
				}).Return(errors.New("error")).Once()
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocker(tt.args, tt.fields)
			c := &CreateCustomerUC{
				CustomerRepository: tt.fields.CustomerRepository,
				UUID:               tt.fields.UUID,
			}
			got, err := c.Execute(tt.args.createCustomerDTO)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateCustomerUC.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateCustomerUC.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
