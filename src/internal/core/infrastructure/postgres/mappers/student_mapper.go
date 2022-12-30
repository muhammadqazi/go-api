package mappers

type StudentMapper interface {
	CreateStudentMapper()
}

type studentMapper struct {
}

func NewStudentMapper() StudentMapper {
	return &studentMapper{}
}

func (m *studentMapper) CreateStudentMapper() {

}
