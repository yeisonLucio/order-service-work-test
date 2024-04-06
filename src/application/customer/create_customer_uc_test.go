package customer

import (
	"errors"
	"reflect"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/mock"
	"lucio.com/order-service/src/domain/common/dtos"
	customerDtos "lucio.com/order-service/src/domain/customer/dtos"
	"lucio.com/order-service/src/domain/customer/entities"
	"lucio.com/order-service/src/mocks"
)

func TestCreateCustomerUC_Execute(t *testing.T) {
	type fields struct {
		CustomerRepository *mocks.CustomerRepository
		Logger             *logrus.Logger
	}
	type args struct {
		createCustomer entities.Customer
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mocker  func(a args, f fields)
		want    *customerDtos.CreatedCustomerResponse
		wantErr bool
	}{
		{
			name: "should create customer successfully",
			fields: fields{
				CustomerRepository: &mocks.CustomerRepository{},
				Logger:             &logrus.Logger{},
			},
			args: args{
				createCustomer: entities.Customer{
					FirstName: "John",
					LastName:  "Doe",
					Address:   "123",
				},
			},
			mocker: func(a args, f fields) {
				customer := entities.Customer{
					FirstName: "John",
					LastName:  "Doe",
					Address:   "123",
				}
				f.CustomerRepository.On("Create", &customer).Return(nil).Once()
			},
			want: &customerDtos.CreatedCustomerResponse{
				IsActive:  false,
				FirstName: "John",
				LastName:  "Doe",
				Address:   "123",
			},
		},
		{
			name: "should return an error when customer cannot be created",
			fields: fields{
				CustomerRepository: &mocks.CustomerRepository{},
				Logger:             &logrus.Logger{},
			},
			args: args{
				createCustomer: entities.Customer{},
			},
			mocker: func(a args, f fields) {
				f.CustomerRepository.On("Create", mock.Anything).Return(&dtos.CustomError{
					Code:  500,
					Error: errors.New("error"),
				}).Once()
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
				Logger:             tt.fields.Logger,
			}
			got, err := c.Execute(tt.args.createCustomer)
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
