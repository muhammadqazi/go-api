package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/api/handlers"
	middleware "github.com/muhammadqazi/SIS-Backend-Go/src/internal/api/middlewares"
)

func CourseRouter(r *gin.RouterGroup, h handlers.CourseHandler) {

	allowedRolesForCreate := []string{"admin"}
	g := r.Group("/course")

	/*
		"""
		We will use the RolesMiddleware to check if the user has the required permissions to access the route
		"""
	*/

	checkRoleForCreate := middleware.RolesMiddleware(allowedRolesForCreate)
	g.Use(checkRoleForCreate)

	g.POST("/create", h.PostCourse)
	g.GET("/get", h.GetCourse)
}
