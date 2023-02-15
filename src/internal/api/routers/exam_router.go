package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/api/handlers"
	middleware "github.com/muhammadqazi/SIS-Backend-Go/src/internal/api/middlewares"
)

func ExamRouter(r *gin.RouterGroup, h handlers.ExamHandler) {

	allowedRolesForCreate := []string{"admin", "instructor"}
	g := r.Group("/exam")

	/*
		"""
		We will use the RolesMiddleware to check if the user has the required permissions to access the route
		"""
	*/

	checkRoleForCreate := middleware.RolesMiddleware(allowedRolesForCreate)
	g.Use(checkRoleForCreate)

	g.POST("/schedule", h.PostExamSchedule)
	g.PATCH("/results", h.PatchExamResults)
}
