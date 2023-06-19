package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadqazi/campus-hq-api/src/internal/common/validation"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/domain/services"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/mappers"
	"net/http"
)

/*
	"""
	RoomHandler can provide the following services.
	"""
*/

type RoomHandler interface {
	PostRoom(c *gin.Context)
	GetRoomByNumber(c *gin.Context)
	GetAvailableRooms(c *gin.Context)
}

type roomHandler struct {
	validator    validation.Validator
	roomMapper   mappers.RoomMapper
	roomServices services.RoomServices
}

/*
	"""
	This will create a new instance of the RoomHandler, we will use this as a constructor
	"""
*/

func NewRoomHandler(service services.RoomServices, mapper mappers.RoomMapper, v validation.Validator) RoomHandler {
	return &roomHandler{
		roomMapper:   mapper,
		roomServices: service,
		validator:    v,
	}
}

func (s *roomHandler) PostRoom(c *gin.Context) {
	var room dtos.RoomCreateDTO

	if err := s.validator.Validate(&room, c); err != nil {
		return
	}

	if err := s.roomServices.CreateRoom(room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Room created successfully"})
}
func (s *roomHandler) GetRoomByNumber(c *gin.Context) {
	number := c.Param("number")

	room, err := s.roomServices.FetchRoomByNumber(number)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "data": room})

}

func (s *roomHandler) GetAvailableRooms(c *gin.Context) {
	rooms, err := s.roomServices.FetchAvailableRooms()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "data": rooms})

}
