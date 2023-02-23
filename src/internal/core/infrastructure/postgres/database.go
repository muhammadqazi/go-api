package core

import (
	"log"

	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Init(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		log.Fatalln(err)
	}

	/*
		"""
		Auto Migrate the database, i.e, create tables automatically if they don't exist.
		Entities are the blueprints of the tables.
		"""
	*/

	db.AutoMigrate(
		&entities.InstructorsEntity{},
		&entities.DepartmentsEntity{},
		&entities.FacultiesEntity{},
		&entities.PersonalInfoEntity{},
		&entities.ContactInfoEntity{},
		&entities.AccountsEntity{},
		&entities.InvoicesEntity{},
		&entities.PaymentsEntity{},
		&entities.StudentsEntity{},
		&entities.CoursesEntity{},
		&entities.CurriculumEntity{},
		&entities.TranscriptEntity{},
		&entities.CourseScheduleEntity{},
		&entities.ExamScheduleEntity{},
		&entities.ExamResultsEntity{},
		&entities.StudentEnrollmentsEntity{},
		&entities.StudentAttendanceEntity{},
		&entities.InstructorEnrollmentsEntity{},
		&entities.CourseCurriculumEntity{},
		&entities.StudentCourseRequestEntity{},
		&entities.InstructorCoursesEntity{},
		&entities.CourseAttendanceEntity{},
		&entities.ExamCourseResultsEntity{},
		&entities.StudentPasswordResetsEntity{},
	)

	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()
}
