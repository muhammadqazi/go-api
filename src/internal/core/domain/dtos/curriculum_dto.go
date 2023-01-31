package dtos

type CurriculumCreateDTO struct {
	OfferSemesters []string `json:"offer_semesters" validate:"required,dive,required"`
	NumberOfYears  int      `json:"number_of_years" validate:"required"`
	DepartmentID   uint     `json:"department_id" validate:"required"`
}
