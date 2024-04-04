package workorder

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
	"lucio.com/order-service/src/domain/common/dtos"
	"lucio.com/order-service/src/domain/workorder/entities"
	"lucio.com/order-service/src/domain/workorder/enums"
)

type CreateWorkOrder struct {
	CustomerID       string `json:"customer_id"`
	Title            string `json:"title"`
	PlannedDateBegin string `json:"planned_date_begin"`
	PlannedDateEnd   string `json:"planned_date_end"`
	Type             string `json:"type" enum:"service_order,inactive_customer"`
}

func (c *CreateWorkOrder) ValidateAndGetEntity(ctx *gin.Context) (*entities.WorkOrder, *dtos.CustomError) {
	incorrectFormatError := &dtos.CustomError{
		Code:  400,
		Error: errors.New("body con formato incorrecto"),
	}

	data, err := validate.FromRequest(ctx.Request)
	if err != nil {
		return nil, incorrectFormatError
	}

	validator := data.Create()

	validator.StringRules(map[string]string{
		"title":              "string|required",
		"planned_date_begin": "string|required",
		"planned_date_end":   "string|required",
		"type":               "string",
	})

	validator.AddMessages(map[string]string{
		"required": "El campo {field} es requerido",
		"string":   "El campo {field} debe ser string",
	})

	if errValidation := validator.ValidateE(); errValidation != nil {
		return nil, &dtos.CustomError{
			Code:  400,
			Error: errors.New(errValidation.String()),
		}
	}

	bodyBytes, _ := json.Marshal(data.Src())
	err = json.Unmarshal(bodyBytes, c)
	if err != nil {
		return nil, incorrectFormatError
	}

	beginDate, err := time.Parse(time.DateTime, c.PlannedDateBegin)
	if err != nil {
		return nil, &dtos.CustomError{
			Code:  400,
			Error: errors.New("el formato del campo planned_date_begin no es correcto"),
		}
	}

	endDate, err := time.Parse(time.DateTime, c.PlannedDateEnd)
	if err != nil {
		return nil, &dtos.CustomError{
			Code:  400,
			Error: errors.New("el formato del campo planned_date_end no es correcto"),
		}
	}

	return &entities.WorkOrder{
		CustomerID:       ctx.Param("id"),
		Title:            c.Title,
		PlannedDateBegin: &beginDate,
		PlannedDateEnd:   &endDate,
		Type:             enums.WorkOrderType(c.Type),
	}, nil
}
