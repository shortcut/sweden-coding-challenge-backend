package validation

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

var errorBag = make(map[string]string)

func Validate(context *gin.Context, s interface{}) (map[string]string, interface{}) {
	
	validate = validator.New()
	
	err := context.BindJSON(&s)

	if err != nil {
		panic(err)
	}

	err = validate.Struct(s)
	
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errorBag[err.Field()] = Map(strings.ToLower(err.Field()), err.Tag())
		}
	}
	
	return errorBag, s
}