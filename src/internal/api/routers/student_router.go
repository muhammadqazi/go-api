package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/api/handlers"
)

func StudentRouter(r *gin.RouterGroup, h handlers.StudentHandler) {

	g := r.Group("/students")

	g.POST("/create", h.CreateStudent)

}
