package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/api/handlers"
)

func InstructorsRouter(r *gin.RouterGroup, h handlers.InstructorsHandler) {

	g := r.Group("/instructors")

	g.POST("/create", h.CreateInstructors)
}
