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
		BindJSON will bind the request body to the struct
		"""
	*/
	if err := c.BindJSON(s); err != nil {
		fmt.Println(err.Error(), "error")
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
		return err
	}

	err := v.validate.Struct(s)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": fmt.Sprintf("%s: %s", e.Namespace(), e.Tag())})
			return err
		}
	}
	return nil
}
