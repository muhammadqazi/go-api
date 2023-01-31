package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/api/handlers"
)

func CurriculumRouter(r *gin.RouterGroup, h handlers.CurriculumHandler) {
	g := r.Group("/curriculum")

	g.POST("/create", h.CreateCurriculum)
}
