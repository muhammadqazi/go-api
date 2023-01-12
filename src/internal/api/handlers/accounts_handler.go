package handlers

import (
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

func (s *accountingHandler) MakePayment(c *gin.Context) {

	var payment dtos.MakePaymentDTO

	id, err := security.IDExtractor(c)
	sid, _ := strconv.ParseUint(id, 10, 64)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": false, "message": err.Error()})
		return
	}

	/*
		"""
		BindJSON will bind the request body to the struct
		"""
	*/

	if err := c.BindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
		return
	}

	if res := s.accountingServices.MakePayment(payment, uint(sid)); res {
		c.JSON(http.StatusOK, gin.H{"status": true, "message": "Trabsaction Sucessfull"})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{"status": true, "message": "An Error occoured while performing this transaction"})

}
