package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	CreateAccounts(c *gin.Context)
}

type accountingHandler struct {
	accountingMapper   mappers.AccountsMapper
	accountingServices services.AccountingServices
}

/*
	"""
	This will creates a new instance of the AccountingHandler, we will use this as a constructor
	"""
*/

func NewAccountingHandler(service services.AccountingServices, mapper mappers.AccountsMapper) AccountingHandler {
	return &accountingHandler{
		accountingMapper:   mapper,
		accountingServices: service,
	}
}

func (s *accountingHandler) CreateAccounts(c *gin.Context) {

	var account dtos.AccountCreateDTO

	/*
		"""
		BindJSON will bind the request body to the struct
		"""
	*/

	if err := c.BindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
		return
	}

	if _, err := s.accountingServices.CreateAccounts(account); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Error creating account"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": false, "message": "Account created sucessfully"})

}
