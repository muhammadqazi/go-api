package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/api/handlers"
	middleware "github.com/muhammadqazi/SIS-Backend-Go/src/internal/api/middlewares"
)

func StudentRouter(r *gin.RouterGroup, h handlers.StudentHandler) {

	allowedRolesForCreate := []string{"admin"}
	allowedRolesForStudent := []string{"student"}

	g := r.Group("/student")
	g.POST("/login", h.PostSignIn)

	/*
		"""
		We will use the RolesMiddleware to check if the user has the required permissions to access the route
		Place all the routes that require authentication below this line
		"""
	*/

	/* Student with Role 'student' */
	checkRoleForStudent := middleware.RolesMiddleware(allowedRolesForStudent)
	g.Use(checkRoleForStudent)
	g.POST("/term/registration", h.PostTermRegistration)
	g.GET("/timetable", h.GetStudentTimetable)
	g.GET("/exam", h.GetStudentExamSchedule)

	/*  Role 'admin' Handlers */
	checkRoleForCreate := middleware.RolesMiddleware(allowedRolesForCreate)
	g.Use(checkRoleForCreate)

	g.POST("/create", h.PostStudent)
	g.GET("/:id", h.GetStudentByID)
}
