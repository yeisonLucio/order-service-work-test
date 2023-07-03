package usecases

import (
	"reflect"
	"testing"

	"lucio.com/order-service/src/dto"
	"lucio.com/order-service/src/helpers"
	"lucio.com/order-service/src/repositories/contracts"
)

func TestCreateWorkOrderUC_Execute(t *testing.T) {
	type fields struct {
		WorkOrderRepository contracts.WorkOrderRepository
		UUID                helpers.UUIDGenerator
		CustomerRepository  contracts.CustomerRepository
		Time                helpers.Timer
	}
	type args struct {
		createWorkOrderDTO dto.CreateWorkOrderDTO
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dto.CreatedWorkOrderDTO
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
