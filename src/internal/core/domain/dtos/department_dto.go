package dtos

type DepartmentCreateDTO struct {
	Name          string `json:"name" validate:"required"`
	Code          string `json:"code" validate:"required"`
	Description   string `json:"description" validate:"required"`
	Dean          string `json:"dean" validate:"required"`
	ViceDean      string `json:"vice_dean" validate:"required"`
	Email         string `json:"email" validate:"required"`
	PhoneNumber   string `json:"phone_number" validate:"required"`
	DeanEmail     string `json:"dean_email" validate:"required"`
	DeanPhone     string `json:"dean_phone" validate:"required"`
	NumberOfYears int    `json:"number_of_years" validate:"required"`

	FacultyID uint `json:"faculty_id" validate:"required"`

	HeadID uint `json:"head_id" validate:"required"`
}
