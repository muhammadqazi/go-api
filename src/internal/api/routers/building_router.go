package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadqazi/campus-hq-api/src/internal/api/handlers"
	middleware "github.com/muhammadqazi/campus-hq-api/src/internal/api/middlewares"
)

func BuildingRouter(r *gin.RouterGroup, h handlers.BuildingHandler) {

	allowedRolesForCreate := []string{"admin"}
	g := r.Group("/building")

	/*
		"""
		We will use the RolesMiddleware to check if the user has the required permissions to access the route
		"""
	*/

	checkRoleForCreate := middleware.RolesMiddleware(allowedRolesForCreate)
	g.Use(checkRoleForCreate)

	g.POST("/create", h.PostBuilding)
	g.GET("/:code", h.GetBuilding)
	g.GET("/all", h.GetAllBuildings)
}
