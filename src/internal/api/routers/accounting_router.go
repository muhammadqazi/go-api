package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/api/handlers"
)

func AccountingRouter(r *gin.RouterGroup, h handlers.AccountingHandler) {

	g := r.Group("/accounts")

	g.POST("/pay", h.MakePayment)
}
