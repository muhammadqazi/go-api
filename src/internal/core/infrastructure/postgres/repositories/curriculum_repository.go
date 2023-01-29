package repositories

import "gorm.io/gorm"

type CurriculumRepository interface {
    InsertCurriculum()
}

type curriculumConnection struct {
    conn *gorm.DB
}

func NewCurriculumRepository(db *gorm.DB) CurriculumRepository {
    return &curriculumConnection{
        conn: db,
    }
}

func (r *curriculumConnection) InsertCurriculum() {

}

