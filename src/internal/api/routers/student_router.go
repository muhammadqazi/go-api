package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadqazi/campus-hq-api/src/internal/api/handlers"
	middleware "github.com/muhammadqazi/campus-hq-api/src/internal/api/middlewares"
)

func StudentRouter(r *gin.RouterGroup, h handlers.StudentHandler) {

	allowedRolesForStudent := []string{"student"}
	allowedRolesForAdmin := []string{"admin"}

	student := r.Group("/student")
	admin := r.Group("/student")

	student.GET("/test", h.TestApi)
	student.POST("/login", h.PostSignIn)
	student.POST("/forgot-password", h.PostForgotPasswordRequest)
	student.PUT("/forgot-password", h.PutForgotPasswordCode)
	student.PATCH("/forgot-password", h.PatchNewPassword)
	student.GET("/:id", h.GetStudentByID)

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
	admin.PATCH("/:id", h.PatchStudent)

	/* Student with Role 'student' */
	checkRoleForStudent := middleware.RolesMiddleware(allowedRolesForStudent)
	student.Use(checkRoleForStudent)
	student.POST("/term/registration", h.PostTermRegistration)
	student.GET("/timetable", h.GetStudentTimetable)
	student.GET("/exam", h.GetStudentExamSchedule)
	student.GET("/attendance", h.GetStudentCoursesAttendance)
	student.PATCH("/reset-password", h.PatchResetPassword)

}
