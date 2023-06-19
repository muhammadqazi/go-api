package dtos

type RoomCreateDTO struct {
	RoomNumber      string `json:"room_number"`
	RoomDescription string `json:"room_description"`
	NumberOfSeats   uint   `json:"number_of_seats"`
	BuildingID      uint   `json:"building_id"`
	IsAvailable     bool   `json:"is_available"`
}
