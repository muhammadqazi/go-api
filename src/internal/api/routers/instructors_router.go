package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/api/handlers"
	middleware "github.com/muhammadqazi/SIS-Backend-Go/src/internal/api/middlewares"
)

func InstructorsRouter(r *gin.RouterGroup, h handlers.InstructorsHandler) {

	allowedRolesForCreate := []string{"admin"}
	allowedRolesForFetch := []string{"admin", "instructor"}

	g := r.Group("/instructor")

	g.POST("/login", h.PostSignIn)

	/*
		"""
		We will use the RolesMiddleware to check if the user has the required permissions to access the route,
		Place all the routes that require authentication below this line
		"""
	*/

	checkRoleForFetch := middleware.RolesMiddleware(allowedRolesForFetch)
	g.Use(checkRoleForFetch)

	g.GET("/requests", h.GetTermEnrollmentRequests)
	g.PATCH("/requests/:request-id", h.PatchTermEnrollmentRequests)
	g.GET("/enrollments/:id", h.GetInstructorCourseEnrollment)
	g.PATCH("/attendance", h.PatchStudentAttendance)

	/* Admin area */
	checkRoleForCreate := middleware.RolesMiddleware(allowedRolesForCreate)
	g.Use(checkRoleForCreate)

	g.POST("/create", h.PostInstructors)
	g.POST("/enrollments", h.PostInstructorCourseEnrollment)
}
