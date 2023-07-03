package models

import (
	"testing"
)

func TestCustomer_Validate(t *testing.T) {
	type fields struct {
		FirstName string
		LastName  string
		Address   string
	}
	tests := []struct {
		name      string
		fields    fields
		wantErr   bool
		errorType error
	}{
		{
			name: "should return an error when first name is empty",
			fields: fields{
				LastName: "Doe",
				Address:  "123",
			},
			wantErr:   true,
			errorType: ErrInvalidCustomerFirstName,
		},
		{
			name: "should return an error when last name is empty",
			fields: fields{
				FirstName: "juan",
				Address:   "123",
			},
			wantErr:   true,
			errorType: ErrInvalidCustomerLastName,
		},
		{
			name: "should return an error when address is empty",
			fields: fields{
				FirstName: "juan",
				LastName:  "Doe",
			},
			wantErr:   true,
			errorType: ErrInvalidCustomerAddress,
		},
		{
			name: "should validate successful",
			fields: fields{
				FirstName: "juan",
				LastName:  "Doe",
				Address:   "123",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Customer{
				FirstName: tt.fields.FirstName,
				LastName:  tt.fields.LastName,
				Address:   tt.fields.Address,
			}
			err := c.Validate()

			if (err != nil) != tt.wantErr {
				t.Errorf("Customer.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}

			if (err != nil) && err != tt.errorType {
				t.Errorf("Customer.Validate() error = %v, errorType %v", err, tt.errorType)
			}
		})
	}
}
