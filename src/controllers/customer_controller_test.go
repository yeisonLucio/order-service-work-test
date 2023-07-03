package controllers

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"lucio.com/order-service/src/dto"
	"lucio.com/order-service/src/mocks"
)

func TestCustomerController_CreateCustomer(t *testing.T) {
	type dependencies struct {
		CreateCustomerUC    *mocks.CreateCustomerUC
		CreateWorkOrderUC   *mocks.CreateWorkOrderUC
		CustomerRepository  *mocks.CustomerRepository
		WorkOrderRepository *mocks.WorkOrderRepository
		UpdateCustomerUC    *mocks.UpdateCustomerUC
	}

	tests := []struct {
		name         string
		dependencies dependencies
		request      []byte
		statusCode   int
		mocker       func(d dependencies)
	}{
		{
			name:       "should return an error when data is invalid",
			request:    []byte(`"first_name":1`),
			statusCode: http.StatusBadRequest,
			mocker:     func(d dependencies) {},
		},
		{
			name: "should return an error when use case fail",
			dependencies: dependencies{
				CreateCustomerUC: &mocks.CreateCustomerUC{},
			},
			request:    []byte(`{}`),
			statusCode: http.StatusInternalServerError,
			mocker: func(d dependencies) {
				d.CreateCustomerUC.On("Execute", dto.CreateCustomerDTO{}).
					Return(&dto.CreatedCustomerDTO{}, errors.New("error")).
					Once()
			},
		},
		{
			name: "should return success when data is ok",
			dependencies: dependencies{
				CreateCustomerUC: &mocks.CreateCustomerUC{},
			},
			request:    []byte(`{"first_name":"juan","last_name":"marin","address":"test"}`),
			statusCode: http.StatusCreated,
			mocker: func(d dependencies) {
				input := dto.CreateCustomerDTO{
					FirstName: "juan",
					LastName:  "marin",
					Address:   "test",
				}

				d.CreateCustomerUC.On("Execute", input).
					Return(nil, nil).
					Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocker(tt.dependencies)
			controller := CustomerController{
				CreateCustomerUC:    tt.dependencies.CreateCustomerUC,
				CreateWorkOrderUC:   tt.dependencies.CreateWorkOrderUC,
				CustomerRepository:  tt.dependencies.CustomerRepository,
				WorkOrderRepository: tt.dependencies.WorkOrderRepository,
				UpdateCustomerUC:    tt.dependencies.UpdateCustomerUC,
			}

			router := gin.Default()
			router.POST("/test", controller.CreateCustomer)
			req, _ := http.NewRequest(http.MethodPost, "/test", bytes.NewBuffer(tt.request))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, req)

			assert.Equal(t, tt.statusCode, rec.Code)
		})
	}
}
