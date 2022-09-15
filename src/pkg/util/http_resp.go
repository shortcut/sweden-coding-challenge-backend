package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MainStructure struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func RespondError(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, MainStructure{
		Meta: Meta{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Success: false,
		},
		Data: nil,
	})
}

func RespondValidationError(c *gin.Context, errs map[string]string) {
	c.JSON(http.StatusUnprocessableEntity, MainStructure{
		Meta: Meta{
			Code:    http.StatusUnprocessableEntity,
			Message: "validation error",
			Success: false,
		},
		Data: errs,
	})
}
