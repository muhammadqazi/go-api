package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadqazi/campus-hq-api/src/internal/api/handlers"
	middleware "github.com/muhammadqazi/campus-hq-api/src/internal/api/middlewares"
)

func RoomRouter(r *gin.RouterGroup, h handlers.RoomHandler) {

	allowedRolesForCreate := []string{"admin"}
	g := r.Group("/room")

	/*
		"""
		We will use the RolesMiddleware to check if the user has the required permissions to access the route
		"""
	*/

	checkRoleForCreate := middleware.RolesMiddleware(allowedRolesForCreate)
	g.Use(checkRoleForCreate)

	g.POST("/create", h.PostRoom)
	g.GET("/:number", h.GetRoomByNumber)
	g.GET("/", h.GetAvailableRooms)
}
