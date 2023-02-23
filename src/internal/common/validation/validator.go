package validation

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type Validator interface {
	Validate(interface{}, *gin.Context) error
}

type structValidator struct {
	validate *validator.Validate
}

func NewStructValidator() Validator {
	return &structValidator{
		validate: validator.New(),
	}
}

func (v *structValidator) Validate(s interface{}, c *gin.Context) error {

	/*
		"""
		BindJSON will bind the request body to the given struct.
		"""
	*/

	if err := c.BindJSON(s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return err
	}

	err := v.validate.Struct(s)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		var errorMessage string

		for _, validationError := range validationErrors {
			fieldName := validationError.Field()
			tag := validationError.Tag()
			param := validationError.Param()

			switch tag {
			case "required":
				errorMessage = fmt.Sprintf("%s is required", fieldName)
			case "min":
				errorMessage = fmt.Sprintf("%s must be at least %s", fieldName, param)
			case "max":
				errorMessage = fmt.Sprintf("%s must be at most %s", fieldName, param)
			default:
				errorMessage = fmt.Sprintf("%s is invalid", fieldName)
			}
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": errorMessage,
		})
		return err
	}

	return nil
}
