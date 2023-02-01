package dtos

type role string

const (
	RootRole       role = "root"
	AdminRole      role = "admin"
	InstructorRole role = "instructor"
	AssistantRole  role = "assistant"
)

type InstructorCreateDTO struct {
	FirstName    string `json:"first_name" validate:"required,min=3,max=255"`
	LastName     string `json:"last_name" validate:"required,min=3,max=255"`
	PhoneNumber  string `json:"phone_number" validate:"required,min=3,max=255"`
	Email        string `json:"email" validate:"required,email"`
	Password     string `json:"password" validate:"required,min=8,max=255"`
	DOB          string `json:"dob" validate:"required,min=3,max=255"`
	PlaceOfBirth string `json:"place_of_birth" validate:"required,min=3,max=255"`
	Sex          string `json:"sex" validate:"required"`
	Nationality  string `json:"nationality" validate:"required`
	Role         role   `json:"role" validate:"required"`
}
