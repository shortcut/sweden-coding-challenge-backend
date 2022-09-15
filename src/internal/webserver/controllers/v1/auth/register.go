package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/mojtabaRKS/link_shortener/pkg/validation"
	"github.com/mojtabaRKS/link_shortener/pkg/util"
)

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email,unique"`
	Password string `json:"password" validate:"required,min=8"`
	Name     string `json:"name" validate:"required"`
	UserName string `json:"username" validate:"required,unique"`
}



func Register(c *gin.Context) {

	errs, request := validation.Validate(c, &RegisterRequest{})

	if len(errs) != 0 {
		util.RespondValidationError(c, errs)
		return
	}

		
}
