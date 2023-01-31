package handlers

import (
	"fmt"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/common/validation"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/common/security"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/domain/services"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/mappers"
)

/*
	"""
	AccountingHandler can provide the following services.
	"""
*/

type AccountingHandler interface {
	MakePayment(c *gin.Context)
}

type accountingHandler struct {
	validator          validation.Validator
	accountingMapper   mappers.AccountsMapper
	accountingServices services.AccountingServices
}

/*
	"""
	This will create a new instance of the AccountingHandler, we will use this as a constructor
	"""
*/

func NewAccountingHandler(service services.AccountingServices, mapper mappers.AccountsMapper, v validation.Validator) AccountingHandler {
	return &accountingHandler{
		accountingMapper:   mapper,
		accountingServices: service,
		validator:          v,
	}
}

func (s *accountingHandler) MakePayment(c *gin.Context) {

	var payment dtos.MakePaymentDTO

	id, err := security.IDExtractor(c)
	sid, _ := strconv.ParseUint(id, 10, 64)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": false, "message": err.Error()})
		return
	}

	if err := s.validator.Validate(&payment, c); err != nil {
		return
	}

	if err := s.accountingServices.MakePayment(payment, uint(sid)); err == nil {
		c.JSON(http.StatusOK, gin.H{"status": true, "message": "Transaction Sucessfull"})
		return
	}

	fmt.Println(err)

	c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "An Error occoured while performing this transaction"})

}
