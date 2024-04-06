package controllers

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	commonDtos "lucio.com/order-service/src/domain/common/dtos"
	"lucio.com/order-service/src/domain/workorder/dtos"

	"lucio.com/order-service/src/domain/workorder/entities"
	"lucio.com/order-service/src/mocks"
)

func TestWorkOrderController_GetWorkOrders(t *testing.T) {
	type dependencies struct {
		WorkOrderRepository *mocks.WorkOrderRepository
		Logger              *logrus.Logger
	}
	tests := []struct {
		name         string
		dependencies dependencies
		parameters   dtos.WorkOrderFilters
		statusCode   int
		mocker       func(d dependencies)
	}{
		{
			name: "should return filters work orders",
			dependencies: dependencies{
				WorkOrderRepository: &mocks.WorkOrderRepository{},
				Logger:              &logrus.Logger{},
			},
			parameters: dtos.WorkOrderFilters{
				PlannedDateBegin: "2023-01-01T00:00:00Z",
				PlannedDateEnd:   "2023-01-01T01:00:00Z",
				Status:           "new",
			},
			statusCode: http.StatusOK,
			mocker: func(d dependencies) {
				input := dtos.WorkOrderFilters{
					PlannedDateBegin: "2023-01-01T00:00:00Z",
					PlannedDateEnd:   "2023-01-01T01:00:00Z",
					Status:           "new",
				}
				d.WorkOrderRepository.
					On("GetFiltered", input).
					Return([]*dtos.WorkOrder{}, nil).
					Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocker(tt.dependencies)
			w := &WorkOrderController{
				WorkOrderRepository: tt.dependencies.WorkOrderRepository,
				Logger:              &logrus.Logger{},
			}

			router := gin.Default()
			router.GET("/work-orders", w.GetWorkOrders)

			queryParams := url.Values{}
			queryParams.Set("planned_date_begin", tt.parameters.PlannedDateBegin)
			queryParams.Set("planned_date_end", tt.parameters.PlannedDateEnd)
			queryParams.Set("status", tt.parameters.Status)

			req, _ := http.NewRequest(
				http.MethodGet,
				"/work-orders?"+queryParams.Encode(),
				nil,
			)

			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, req)

			assert.Equal(t, tt.statusCode, rec.Code)
		})
	}
}

func TestWorkOrderController_FinishWorkOrder(t *testing.T) {
	type dependencies struct {
		FinishWorkOrderUC *mocks.FinishWorkOrderUC
		Logger            *logrus.Logger
	}
	tests := []struct {
		name         string
		dependencies dependencies
		workOrderId  string
		statusCode   int
		mocker       func(d dependencies)
	}{
		{
			name: "should finish work order",
			dependencies: dependencies{
				FinishWorkOrderUC: &mocks.FinishWorkOrderUC{},
				Logger:            &logrus.Logger{},
			},
			workOrderId: "uuid",
			statusCode:  http.StatusNoContent,
			mocker: func(d dependencies) {
				d.FinishWorkOrderUC.
					On("Execute", "uuid").
					Return(nil).
					Once()
			},
		},
		{
			name: "should return an error when order cannot be finished",
			dependencies: dependencies{
				FinishWorkOrderUC: &mocks.FinishWorkOrderUC{},
				Logger:            &logrus.Logger{},
			},
			workOrderId: "uuid",
			statusCode:  http.StatusInternalServerError,
			mocker: func(d dependencies) {
				d.FinishWorkOrderUC.
					On("Execute", "uuid").
					Return(errors.New("error")).
					Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocker(tt.dependencies)
			w := &WorkOrderController{
				FinishWorkOrderUC: tt.dependencies.FinishWorkOrderUC,
				Logger:            &logrus.Logger{},
			}
			router := gin.Default()
			router.PATCH("/work-orders/:id/finish", w.FinishWorkOrder)
			path := fmt.Sprintf("/work-orders/%s/finish", tt.workOrderId)
			req, _ := http.NewRequest(
				http.MethodPatch,
				path,
				nil,
			)
			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, req)
			assert.Equal(t, tt.statusCode, rec.Code)
		})
	}
}

func TestWorkOrderController_GetWorkOrder(t *testing.T) {
	type dependencies struct {
		WorkOrderRepository *mocks.WorkOrderRepository
	}
	tests := []struct {
		name         string
		dependencies dependencies
		WorkOrderId  string
		statusCode   int
		mocker       func(d dependencies)
	}{
		{
			name: "should get work order successfully",
			dependencies: dependencies{
				WorkOrderRepository: &mocks.WorkOrderRepository{},
			},
			WorkOrderId: "uuid",
			statusCode:  http.StatusOK,
			mocker: func(d dependencies) {
				filters := dtos.WorkOrderFilters{
					ID: "uuid",
				}
				d.WorkOrderRepository.On("GetFiltered", filters).
					Return([]*dtos.WorkOrder{
						{
							ID: "12",
						},
					}).
					Once()
			},
		},
		{
			name: "should return an error when work order does not exist",
			dependencies: dependencies{
				WorkOrderRepository: &mocks.WorkOrderRepository{},
			},
			statusCode: http.StatusNotFound,
			mocker: func(d dependencies) {
				d.WorkOrderRepository.On("GetFiltered", mock.Anything).
					Return([]*dtos.WorkOrder{}).
					Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocker(tt.dependencies)
			w := &WorkOrderController{
				WorkOrderRepository: tt.dependencies.WorkOrderRepository,
			}
			router := gin.Default()
			router.GET("/work-orders/:id", w.GetWorkOrder)
			path := fmt.Sprintf("/work-orders/%s", tt.WorkOrderId)
			req, _ := http.NewRequest(
				http.MethodGet,
				path,
				nil,
			)
			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, req)
			assert.Equal(t, tt.statusCode, rec.Code)

		})
	}
}

func TestWorkOrderController_UpdateWorkOrder(t *testing.T) {
	type dependencies struct {
		UpdateWorkOrderUC *mocks.UpdateWorkOrderUC
		Logger            *logrus.Logger
	}

	plannedDate, _ := time.Parse(time.DateTime, "2024-03-29 19:14:00")

	tests := []struct {
		name         string
		dependencies dependencies
		workOrderId  string
		body         []byte
		statusCode   int
		mocker       func(d dependencies)
	}{
		{
			name: "should update work order successfully",
			dependencies: dependencies{
				UpdateWorkOrderUC: &mocks.UpdateWorkOrderUC{},
				Logger:            &logrus.Logger{},
			},
			workOrderId: "uuid",
			body: []byte(`{
				"title":"test",
				"planned_date_begin":"2024-03-29 19:14:00",
				"planned_date_end":"2024-03-29 19:14:00",
				"type":"test"
				}`),
			statusCode: http.StatusOK,
			mocker: func(d dependencies) {
				input := entities.WorkOrder{
					ID:               "uuid",
					Title:            "test",
					PlannedDateBegin: &plannedDate,
					PlannedDateEnd:   &plannedDate,
					Type:             "test",
				}
				d.UpdateWorkOrderUC.
					On("Execute", input).
					Return(nil, nil).
					Once()
			},
		},
		{
			name:        "should return an error when data is not valid",
			workOrderId: "uuid",
			body:        []byte(`{"title":1}`),
			dependencies: dependencies{
				Logger: &logrus.Logger{},
			},
			statusCode: http.StatusBadRequest,
			mocker:     func(d dependencies) {},
		},
		{
			name: "should return an error when work order cannot be updated",
			dependencies: dependencies{
				UpdateWorkOrderUC: &mocks.UpdateWorkOrderUC{},
				Logger:            &logrus.Logger{},
			},
			workOrderId: "uuid",
			body:        []byte(`{}`),
			statusCode:  http.StatusInternalServerError,
			mocker: func(d dependencies) {
				d.UpdateWorkOrderUC.
					On("Execute", entities.WorkOrder{}).
					Return(nil, errors.New("error")).
					Once()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocker(tt.dependencies)
			w := &WorkOrderController{
				UpdateWorkOrderUC: tt.dependencies.UpdateWorkOrderUC,
				Logger:            tt.dependencies.Logger,
			}
			router := gin.Default()
			router.PUT("/work-orders/:id", w.UpdateWorkOrder)
			path := fmt.Sprintf("/work-orders/%s", tt.workOrderId)
			req, _ := http.NewRequest(
				http.MethodPut,
				path,
				bytes.NewBuffer(tt.body),
			)
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, req)
			assert.Equal(t, tt.statusCode, rec.Code)
		})
	}
}

func TestWorkOrderController_DeleteWorkOrder(t *testing.T) {
	type dependencies struct {
		WorkOrderRepository *mocks.WorkOrderRepository
		Logger              *logrus.Logger
	}
	tests := []struct {
		name         string
		dependencies dependencies
		customerId   string
		statusCode   int
		mocker       func(d dependencies)
	}{
		{
			name: "should return an error when work order does not exist",
			dependencies: dependencies{
				WorkOrderRepository: &mocks.WorkOrderRepository{},
				Logger:              &logrus.Logger{},
			},
			customerId: "123",
			statusCode: http.StatusNotFound,
			mocker: func(d dependencies) {
				d.WorkOrderRepository.On("DeleteByID", "123").Return(&commonDtos.CustomError{
					Code:  404,
					Error: errors.New("error"),
				}).Once()
			},
		},
		{
			name: "should delete the work order successfully",
			dependencies: dependencies{
				WorkOrderRepository: &mocks.WorkOrderRepository{},
				Logger:              &logrus.Logger{},
			},
			customerId: "123",
			statusCode: http.StatusNoContent,
			mocker: func(d dependencies) {
				d.WorkOrderRepository.On("DeleteByID", "123").Return(nil).Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocker(tt.dependencies)
			w := WorkOrderController{
				WorkOrderRepository: tt.dependencies.WorkOrderRepository,
				Logger:              tt.dependencies.Logger,
			}
			router := gin.Default()
			router.DELETE("/work-orders/:id", w.DeleteWorkOrder)
			path := fmt.Sprintf("/work-orders/%s", tt.customerId)
			req, _ := http.NewRequest(http.MethodDelete, path, nil)
			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, req)
			assert.Equal(t, tt.statusCode, rec.Code)
		})
	}
}
