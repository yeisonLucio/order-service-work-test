package customer

import (
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
	"lucio.com/order-service/src/domain/common/dtos"
	customerDtos "lucio.com/order-service/src/domain/customer/dtos"
	"lucio.com/order-service/src/domain/customer/entities"
	"lucio.com/order-service/src/mocks"
)

func TestUpdateCustomerUC_Execute(t *testing.T) {
	type fields struct {
		CustomerRepository *mocks.CustomerRepository
	}
	type args struct {
		updateCustomer entities.Customer
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mocker  func(a args, f fields)
		want    *customerDtos.UpdatedCustomerResponse
		wantErr bool
	}{
		{
			name: "should return an error when customer does not exist",
			fields: fields{
				CustomerRepository: &mocks.CustomerRepository{},
			},
			args: args{
				updateCustomer: entities.Customer{},
			},
			mocker: func(a args, f fields) {
				f.CustomerRepository.On(
					"FindByID",
					a.updateCustomer.ID,
				).Return(nil, &dtos.CustomError{
					Code:  404,
					Error: errors.New("error"),
				}).Once()
			},
			wantErr: true,
		},
		{
			name: "should return an error when customer cannot be saved",
			fields: fields{
				CustomerRepository: &mocks.CustomerRepository{},
			},
			args: args{
				updateCustomer: entities.Customer{
					FirstName: "John",
					LastName:  "Doe",
					Address:   "fake",
				},
			},
			mocker: func(a args, f fields) {
				f.CustomerRepository.On(
					"FindByID",
					a.updateCustomer.ID,
				).Return(&entities.Customer{}, nil).Once()

				f.CustomerRepository.On(
					"Save",
					mock.Anything,
				).Return(&dtos.CustomError{
					Code:  500,
					Error: errors.New("error"),
				}).Once()
			},
			wantErr: true,
		},
		{
			name: "should update customer successfully",
			fields: fields{
				CustomerRepository: &mocks.CustomerRepository{},
			},
			args: args{
				updateCustomer: entities.Customer{
					FirstName: "juan",
					LastName:  "Doe",
					Address:   "fake",
				},
			},
			mocker: func(a args, f fields) {
				customer := entities.Customer{
					ID:        "abc12",
					FirstName: "john",
					LastName:  "Doe",
					Address:   "fake",
				}
				f.CustomerRepository.On(
					"FindByID",
					a.updateCustomer.ID,
				).Return(&customer, nil).Once()

				customer.FirstName = "juan"

				f.CustomerRepository.On(
					"Save",
					&customer,
				).Return(nil).Once()
			},
			want: &customerDtos.UpdatedCustomerResponse{
				ID:        "abc12",
				FirstName: "juan",
				LastName:  "Doe",
				Address:   "fake",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocker(tt.args, tt.fields)
			u := &UpdateCustomerUC{
				CustomerRepository: tt.fields.CustomerRepository,
			}
			got, err := u.Execute(tt.args.updateCustomer)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateCustomerUC.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateCustomerUC.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
