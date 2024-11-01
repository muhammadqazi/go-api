package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadqazi/campus-hq-api/src/internal/api/handlers"
	middleware "github.com/muhammadqazi/campus-hq-api/src/internal/api/middlewares"
)

func AccountingRouter(r *gin.RouterGroup, h handlers.AccountingHandler) {

	allowedRolesForStudent := []string{"student"}
	allowedRolesForAdmin := []string{"admin"}

	student := r.Group("/accounts")
	admin := r.Group("/accounts")

	student.GET("/:id", h.GetAccountDetailsByStudentId)

	/*
		"""
		We will use the RolesMiddleware to check if the user has the required permissions to access the route
		Place all the routes that require authentication below this line
		"""
	*/

	/*  Role 'student' Handlers */
	checkRoleForStudent := middleware.RolesMiddleware(allowedRolesForStudent)
	student.Use(checkRoleForStudent)

	student.POST("/pay", h.PostPayment)

	/* Handlers with 'admin' roles */

	admin.Use(middleware.RolesMiddleware(allowedRolesForAdmin))

	//admin.PATCH("/info", h.PatchAccountDetailsByStudentID)
}
