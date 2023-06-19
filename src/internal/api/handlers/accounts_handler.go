package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/muhammadqazi/campus-hq-api/src/internal/common/validation"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/domain/services"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/mappers"
	"net/http"
	"strconv"
)

/*
	"""
	AccountingHandler can provide the following services.
	"""
*/

type AccountingHandler interface {
	PostPayment(c *gin.Context)
	GetAccountDetailsByStudentId(c *gin.Context)
	PatchAccountDetailsByStudentID(c *gin.Context)
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

func (s *accountingHandler) PostPayment(c *gin.Context) {

	var payment dtos.MakePaymentDTO

	//id := c.MustGet("id").(string)
	//sid, _ := strconv.ParseUint(id, 10, 64)

	if err := s.validator.Validate(&payment, c); err != nil {
		return
	}

	if err := s.accountingServices.CreatePayment(payment, uint(payment.StudentID)); err == nil {
		c.JSON(http.StatusOK, gin.H{"status": true, "message": "Transaction Successful"})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "An Error occurred while performing this transaction"})

}

func (s *accountingHandler) GetAccountDetailsByStudentId(c *gin.Context) {

	id := c.Param("id")
	sid, _ := strconv.ParseUint(id, 10, 64)

	fmt.Println(sid, "we here")

	if account, err := s.accountingServices.FetchAccountDetails(uint(sid)); err == nil {

		if len(account) == 0 {
			c.JSON(http.StatusOK, gin.H{"status": true, "message": "No Account Found"})
			return
		}

		mappedData := s.accountingMapper.AccountsFetchMapper(account)
		c.JSON(http.StatusOK, gin.H{"status": true, "data": mappedData})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "An Error occurred while performing this transaction"})

}

func (s *accountingHandler) PatchAccountDetailsByStudentID(c *gin.Context) {
}
