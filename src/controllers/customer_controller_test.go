package controllers

import (
	"bytes"
	"errors"
	"fmt"
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
		CreateCustomerUC *mocks.CreateCustomerUC
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
			c := CustomerController{
				CreateCustomerUC: tt.dependencies.CreateCustomerUC,
			}

			router := gin.Default()
			router.POST("/test", c.CreateCustomer)
			req, _ := http.NewRequest(http.MethodPost, "/test", bytes.NewBuffer(tt.request))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, req)

			assert.Equal(t, tt.statusCode, rec.Code)
		})
	}
}

func TestCustomerController_CreateWorkOrder(t *testing.T) {
	type dependencies struct {
		CreateWorkOrderUC *mocks.CreateWorkOrderUC
	}
	tests := []struct {
		name         string
		dependencies dependencies
		customerId   string
		body         []byte
		statusCode   int
		mocker       func(d dependencies)
	}{
		{
			name:       "should return an error when data is invalid",
			customerId: "123",
			body:       []byte(`"title":1`),
			statusCode: http.StatusBadRequest,
			mocker:     func(d dependencies) {},
		},
		{
			name: "should return an error when use case fail",
			dependencies: dependencies{
				CreateWorkOrderUC: &mocks.CreateWorkOrderUC{},
			},
			customerId: "123",
			body:       []byte(`{}`),
			statusCode: http.StatusInternalServerError,
			mocker: func(d dependencies) {
				d.CreateWorkOrderUC.On("Execute", dto.CreateWorkOrderDTO{}).
					Return(&dto.CreatedCustomerDTO{}, errors.New("error")).
					Once()
			},
		},
		{
			name: "should return success when data is ok",
			dependencies: dependencies{
				CreateWorkOrderUC: &mocks.CreateWorkOrderUC{},
			},
			customerId: "123",
			body: []byte(`{
				"title":"test",
				"planned_date_begin":"test",
				"planned_date_end":"test",
				"type":"test"
				}`),
			statusCode: http.StatusCreated,
			mocker: func(d dependencies) {
				input := dto.CreateWorkOrderDTO{
					CustomerID:       "123",
					Title:            "test",
					PlannedDateBegin: "test",
					PlannedDateEnd:   "test",
					Type:             "test",
				}

				d.CreateWorkOrderUC.On("Execute", input).
					Return(nil, nil).
					Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocker(tt.dependencies)
			c := &CustomerController{
				CreateWorkOrderUC: tt.dependencies.CreateWorkOrderUC,
			}

			router := gin.Default()
			router.POST("/test/:id", c.CreateWorkOrder)
			path := fmt.Sprintf("/test/%s", tt.customerId)
			req, _ := http.NewRequest(http.MethodPost, path, bytes.NewBuffer(tt.body))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, req)

			assert.Equal(t, tt.statusCode, rec.Code)
		})
	}
}

func TestCustomerController_GetWorkOrders(t *testing.T) {
	type dependencies struct {
		WorkOrderRepository *mocks.WorkOrderRepository
	}

	tests := []struct {
		name         string
		dependencies dependencies
		customerId   string
		statusCode   int
		mocker       func(d dependencies)
	}{
		{
			name: "should return work orders according to customer",
			dependencies: dependencies{
				WorkOrderRepository: &mocks.WorkOrderRepository{},
			},
			customerId: "123",
			statusCode: http.StatusOK,
			mocker: func(d dependencies) {
				d.WorkOrderRepository.On("GetByCustomerID", "123").
					Return(nil, nil).
					Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocker(tt.dependencies)
			c := CustomerController{
				WorkOrderRepository: tt.dependencies.WorkOrderRepository,
			}

			router := gin.Default()
			router.GET("/test/:id", c.GetWorkOrders)
			path := fmt.Sprintf("/test/%s", tt.customerId)
			req, _ := http.NewRequest(http.MethodGet, path, nil)
			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, req)

			assert.Equal(t, tt.statusCode, rec.Code)
		})
	}
}

func TestCustomerController_GetCustomers(t *testing.T) {
	type dependencies struct {
		CustomerRepository *mocks.CustomerRepository
	}
	tests := []struct {
		name         string
		dependencies dependencies
		statusCode   int
		mocker       func(d dependencies)
	}{
		{
			name: "should return active customers",
			dependencies: dependencies{
				CustomerRepository: &mocks.CustomerRepository{},
			},
			statusCode: http.StatusOK,
			mocker: func(d dependencies) {
				d.CustomerRepository.On("GetActives").
					Return(nil, nil).
					Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocker(tt.dependencies)
			c := CustomerController{
				CustomerRepository: tt.dependencies.CustomerRepository,
			}

			router := gin.Default()
			router.GET("/test", c.GetCustomers)
			req, _ := http.NewRequest(http.MethodGet, "/test", nil)
			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, req)

			assert.Equal(t, tt.statusCode, rec.Code)
		})
	}
}

func TestCustomerController_GetCustomer(t *testing.T) {
	type dependencies struct {
		CustomerRepository *mocks.CustomerRepository
	}
	tests := []struct {
		name         string
		dependencies dependencies
		customerId   string
		statusCode   int
		mocker       func(d dependencies)
	}{
		{
			name: "should return a customer according to id",
			dependencies: dependencies{
				CustomerRepository: &mocks.CustomerRepository{},
			},
			customerId: "123",
			statusCode: http.StatusOK,
			mocker: func(d dependencies) {
				d.CustomerRepository.On("GetByID", "123").
					Return(nil, nil).
					Once()
			},
		},
		{
			name: "should return an error when customer does not exist",
			dependencies: dependencies{
				CustomerRepository: &mocks.CustomerRepository{},
			},
			customerId: "123",
			statusCode: http.StatusNotFound,
			mocker: func(d dependencies) {
				d.CustomerRepository.On("GetByID", "123").Return(nil, errors.New("error")).Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocker(tt.dependencies)
			c := CustomerController{
				CustomerRepository: tt.dependencies.CustomerRepository,
			}

			router := gin.Default()
			router.GET("/test/:id", c.GetCustomer)
			path := fmt.Sprintf("/test/%s", tt.customerId)
			req, _ := http.NewRequest(http.MethodGet, path, nil)
			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, req)

			assert.Equal(t, tt.statusCode, rec.Code)
		})
	}

}

func TestCustomerController_UpdateCustomer(t *testing.T) {
	type dependencies struct {
		UpdateCustomerUC *mocks.UpdateCustomerUC
	}
	tests := []struct {
		name         string
		dependencies dependencies
		customerId   string
		body         []byte
		statusCode   int
		mocker       func(d dependencies)
	}{
		{
			name:       "should return an error when data is invalid",
			customerId: "123",
			body:       []byte(`"first_name":1`),
			statusCode: http.StatusBadRequest,
			mocker:     func(d dependencies) {},
		},
		{
			name: "should return success when data is valid",
			dependencies: dependencies{
				UpdateCustomerUC: &mocks.UpdateCustomerUC{},
			},
			customerId: "123",
			body:       []byte(`{"first_name":"test","last_name":"test","address":"test"}`),
			statusCode: http.StatusOK,
			mocker: func(d dependencies) {
				input := dto.UpdateCustomerDTO{
					ID:        "123",
					FirstName: "test",
					LastName:  "test",
					Address:   "test",
				}
				d.UpdateCustomerUC.On("Execute", input).Return(nil, nil).Once()
			},
		},
		{
			name: "should return an error when the use case fail",
			dependencies: dependencies{
				UpdateCustomerUC: &mocks.UpdateCustomerUC{},
			},
			customerId: "123",
			body:       []byte(`{}`),
			statusCode: http.StatusInternalServerError,
			mocker: func(d dependencies) {
				input := dto.UpdateCustomerDTO{
					ID: "123",
				}

				d.UpdateCustomerUC.On("Execute", input).
					Return(nil, errors.New("error")).
					Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocker(tt.dependencies)
			c := CustomerController{
				UpdateCustomerUC: tt.dependencies.UpdateCustomerUC,
			}
			router := gin.Default()
			router.PUT("/test/:id", c.UpdateCustomer)
			path := fmt.Sprintf("/test/%s", tt.customerId)
			req, _ := http.NewRequest(http.MethodPut, path, bytes.NewBuffer(tt.body))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, req)
			assert.Equal(t, tt.statusCode, rec.Code)
		})
	}
}

func TestCustomerController_DeleteCustomer(t *testing.T) {
	type dependencies struct {
		CustomerRepository *mocks.CustomerRepository
	}
	tests := []struct {
		name         string
		dependencies dependencies
		customerId   string
		statusCode   int
		mocker       func(d dependencies)
	}{
		{
			name: "should return an error when customer does not exist",
			dependencies: dependencies{
				CustomerRepository: &mocks.CustomerRepository{},
			},
			customerId: "123",
			statusCode: http.StatusNotFound,
			mocker: func(d dependencies) {
				d.CustomerRepository.On("DeleteByID", "123").Return(errors.New("error")).Once()
			},
		},
		{
			name: "should delete customer successfully",
			dependencies: dependencies{
				CustomerRepository: &mocks.CustomerRepository{},
			},
			customerId: "123",
			statusCode: http.StatusNoContent,
			mocker: func(d dependencies) {
				d.CustomerRepository.On("DeleteByID", "123").Return(nil).Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocker(tt.dependencies)
			c := CustomerController{
				CustomerRepository: tt.dependencies.CustomerRepository,
			}
			router := gin.Default()
			router.DELETE("/test/:id", c.DeleteCustomer)
			path := fmt.Sprintf("/test/%s", tt.customerId)
			req, _ := http.NewRequest(http.MethodDelete, path, nil)
			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, req)
			assert.Equal(t, tt.statusCode, rec.Code)
		})
	}
}
