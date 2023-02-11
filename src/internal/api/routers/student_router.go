package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/api/handlers"
	middleware "github.com/muhammadqazi/SIS-Backend-Go/src/internal/api/middlewares"
)

func StudentRouter(r *gin.RouterGroup, h handlers.StudentHandler) {

	allowedRolesForCreate := []string{"admin"}

	g := r.Group("/student")

	g.POST("/login", h.PostStudentSignIn)
	g.POST("/term/registration", h.PostTermRegistration)

	/*
		"""
		We will use the RolesMiddleware to check if the user has the required permissions to access the route
		Place all the routes that require authentication below this line
		"""
	*/

	checkRoleForCreate := middleware.RolesMiddleware(allowedRolesForCreate)
	g.Use(checkRoleForCreate)

	g.POST("/create", h.PostCreateStudent)
	g.GET("/:id", h.GetStudentByID)
}
