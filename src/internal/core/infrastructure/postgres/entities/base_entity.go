package entities

import "time"

type BaseEntity struct {
	CreatedAt time.Time `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updated_at"`
	DeletedAt time.Time `gorm:"type:timestamp" json:"deleted_at"`
	IsActive  bool      `gorm:"type:bool;not null" json:"is_active"`
}

// INSERT INTO students_entity (created_at, updated_at, deleted_at, is_active, first_name, surname, email, nationality, dob, place_of_birth, sex, password, role, status, semester, enrollment_date, graduation_date,student_id)
// VALUES (current_timestamp, current_timestamp, NULL, true, 'John', 'Doe', 'john.doe@example.com', 'American', '01/01/1970', 'New York', 'male', 'password', 'student', 'enrolled', 'Spring', '01/01/2021', NULL,21906778),
//        (current_timestamp, current_timestamp, NULL, true, 'Jane', 'Doe', 'jane.doe@example.com', 'Canadian', '01/01/1980', 'Toronto', 'female', 'password', 'student', 'enrolled', 'Fall', '01/01/2021', NULL,22107446);
