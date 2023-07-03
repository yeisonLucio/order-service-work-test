package usecases

import (
	"errors"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"lucio.com/order-service/src/dto"
	"lucio.com/order-service/src/mocks"
	"lucio.com/order-service/src/models"
)

func TestUpdateCustomerUC_Execute(t *testing.T) {
	type fields struct {
		CustomerRepository *mocks.CustomerRepository
	}
	type args struct {
		updateCustomerDTO dto.UpdateCustomerDTO
	}
	uuid := uuid.New()
	tests := []struct {
		name    string
		fields  fields
		args    args
		mocker  func(a args, f fields)
		want    *dto.CustomerDTO
		wantErr bool
	}{
		{
			name: "should return an error when customer does not exist",
			fields: fields{
				CustomerRepository: &mocks.CustomerRepository{},
			},
			args: args{
				updateCustomerDTO: dto.UpdateCustomerDTO{},
			},
			mocker: func(a args, f fields) {
				f.CustomerRepository.On(
					"FindByID",
					a.updateCustomerDTO.ID,
				).Return(nil, errors.New("error")).Once()
			},
			wantErr: true,
		},
		{
			name: "should return an error when customer cannot be saved",
			fields: fields{
				CustomerRepository: &mocks.CustomerRepository{},
			},
			args: args{
				updateCustomerDTO: dto.UpdateCustomerDTO{
					FirstName: "John",
					LastName:  "Doe",
					Address:   "fake",
				},
			},
			mocker: func(a args, f fields) {
				f.CustomerRepository.On(
					"FindByID",
					a.updateCustomerDTO.ID,
				).Return(&models.Customer{}, nil).Once()

				f.CustomerRepository.On(
					"Save",
					mock.Anything,
				).Return(errors.New("error")).Once()
			},
			wantErr: true,
		},
		{
			name: "should update customer successfully",
			fields: fields{
				CustomerRepository: &mocks.CustomerRepository{},
			},
			args: args{
				updateCustomerDTO: dto.UpdateCustomerDTO{
					FirstName: "juan",
					LastName:  "Doe",
					Address:   "fake",
				},
			},
			mocker: func(a args, f fields) {
				customer := models.Customer{
					ID:        uuid,
					FirstName: "john",
					LastName:  "Doe",
					Address:   "fake",
				}
				f.CustomerRepository.On(
					"FindByID",
					a.updateCustomerDTO.ID,
				).Return(&customer, nil).Once()

				customer.FirstName = "juan"

				f.CustomerRepository.On(
					"Save",
					&customer,
				).Return(nil).Once()
			},
			want: &dto.CustomerDTO{
				ID:        uuid.String(),
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
			got, err := u.Execute(tt.args.updateCustomerDTO)
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
