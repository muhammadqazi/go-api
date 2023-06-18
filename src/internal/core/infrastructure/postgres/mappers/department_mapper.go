package mappers

type DepartmentMapper interface {
    DepartmentCreateMapper()
}

type departmentMapper struct {
}

func NewDepartmentMapper() DepartmentMapper {
    return &departmentMapper{}
}

func (m *departmentMapper) DepartmentCreateMapper() {
}

