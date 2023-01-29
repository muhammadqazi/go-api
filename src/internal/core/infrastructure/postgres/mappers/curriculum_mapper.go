package mappers

type CurriculumMapper interface {
}

type curriculumMapper struct {
}

func NewCurriculumMapper() CurriculumMapper {
    return &curriculumMapper{}
}

