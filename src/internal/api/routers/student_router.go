package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/api/handlers"
	middleware "github.com/muhammadqazi/SIS-Backend-Go/src/internal/api/middlewares"
)

func StudentRouter(r *gin.RouterGroup, h handlers.StudentHandler) {

	allowedRolesForStudent := []string{"student"}
	allowedRolesForAdmin := []string{"admin"}

	student := r.Group("/student")
	admin := r.Group("/student")

	student.POST("/login", h.PostSignIn)

	/*
		"""
		We will use the RolesMiddleware to check if the user has the required permissions to access the route
		Place all the routes that require authentication below this line
		"""
	*/

	/*  Role 'admin' Handlers */
	checkRoleForCreate := middleware.RolesMiddleware(allowedRolesForAdmin)
	admin.Use(checkRoleForCreate)

	admin.POST("/create", h.PostStudent)
	admin.GET("/:id", h.GetStudentByID)

	/* Student with Role 'student' */
	checkRoleForStudent := middleware.RolesMiddleware(allowedRolesForStudent)
	student.Use(checkRoleForStudent)

	student.POST("/term/registration", h.PostTermRegistration)
	student.GET("/timetable", h.GetStudentTimetable)
	student.GET("/exam", h.GetStudentExamSchedule)
	student.GET("/attendance", h.GetStudentAttendance)

}
