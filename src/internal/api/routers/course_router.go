package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadqazi/campus-hq-api/src/internal/api/handlers"
	middleware "github.com/muhammadqazi/campus-hq-api/src/internal/api/middlewares"
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
	g.GET("/:code", h.GetCourseByCourseCode)
	g.PATCH("/:code", h.PatchCourseByCourseCode)
	g.DELETE("/:code", h.DeleteCourseByCourseCode)

}
