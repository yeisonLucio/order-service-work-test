package customer

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
	"lucio.com/order-service/src/domain/common/dtos"
	"lucio.com/order-service/src/domain/customer/entities"
)

// UpdateCustomer define el request para actualizar un customer
type UpdateCustomer struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
}

// ValidateAndGetEntity permite validar el request de actualizar un customer
func (c *UpdateCustomer) ValidateAndGetEntity(ctx *gin.Context) (*entities.Customer, *dtos.CustomError) {
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
		"first_name": "string",
		"last_name":  "string",
		"address":    "string",
	})

	validator.AddMessages(map[string]string{
		"string": "El campo {field} debe ser string",
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

	return &entities.Customer{
		ID:        ctx.Param("id"),
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Address:   c.Address,
	}, nil
}
