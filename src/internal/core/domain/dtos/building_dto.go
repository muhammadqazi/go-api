package dtos

type BuildingCreateDTO struct {
	Name          string `json:"name" binding:"required"`
	Code          string `json:"code" binding:"required"`
	Description   string `json:"description" binding:"required"`
	NumberOfRooms int    `json:"number_of_rooms" binding:"required"`
}
