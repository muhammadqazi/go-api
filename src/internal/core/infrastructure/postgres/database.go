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
		&entities.StudentsEntity{},
		&entities.StdPaymentsEntity{},
		&entities.StdInvoicesEntity{},
		&entities.StdAccountingInfoEntity{},
		&entities.PersonalInfoEntity{},
		&entities.FacultiesEntity{},
		&entities.DepartmentsEntity{},
		&entities.CoursesEntity{},
		&entities.ContactInfoEntity{},
		&entities.ContactInfoEntity{},
		&entities.AddressesEntity{},
		&entities.AddressesEntity{},
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
