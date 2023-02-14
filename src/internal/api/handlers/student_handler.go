package handlers

import (
	"errors"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/common/utils"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/common/validation"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/common/security"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/domain/services"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/mappers"
	"gorm.io/gorm"
)

/*
	"""
	StudentHandler can provide the following services.
	"""
*/

type StudentHandler interface {
	PostStudent(c *gin.Context)
	GetStudentByEmail(c *gin.Context)
	GetStudentByID(c *gin.Context)
	PostSignIn(c *gin.Context)
	PostTermRegistration(c *gin.Context)
	GetStudentTimetable(c *gin.Context)
	GetStudentExamSchedule(c *gin.Context)
}

type studentHandler struct {
	validator       validation.Validator
	studentMapper   mappers.StudentMapper
	studentServices services.StudentServices
	accountServices services.AccountingServices
	jwtService      security.TokenManager
}

/*
	"""
	This will create a new instance of the StudentHandler, we will use this as a constructor
	"""
*/

func NewStudentsHandler(service services.StudentServices, account services.AccountingServices, mapper mappers.StudentMapper, jwtService security.TokenManager, v validation.Validator) StudentHandler {
	return &studentHandler{
		studentMapper:   mapper,
		studentServices: service,
		accountServices: account,
		jwtService:      jwtService,
		validator:       v,
	}
}

func (s *studentHandler) PostStudent(c *gin.Context) {

	var student dtos.StudentCreateDTO

	if err := s.validator.Validate(&student, c); err != nil {
		return
	}

	/*
		"""
		We will check if the student already exists in the database
		"""
	*/

	_, err := s.studentServices.FetchStudentByEmail(student.Email)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Student already exists"})
		return
	}

	recentSid, err := s.studentServices.FetchLastStudentID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Error generating student number"})
		return
	}

	semester := utils.GetCurrentSemester()
	sid := recentSid + 1

	if sid, err := s.studentServices.CreateStudent(student, sid, semester); err == nil {
		if _, err := s.accountServices.CreateAccounts(student, sid); err == nil {
			c.JSON(http.StatusOK, gin.H{"status": true, "message": "Student created successfully", "student_id": sid})
			return
		}
	}

	c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Error creating student"})

}

func (s *studentHandler) GetStudentByEmail(c *gin.Context) {

}

func (s *studentHandler) PostSignIn(c *gin.Context) {

	var student dtos.StudentSignInDTO

	if err := s.validator.Validate(&student, c); err != nil {
		return
	}

	if doc, err := s.studentServices.FetchStudentByID(student.StudentID); err == nil {

		password := security.CheckPasswordHash(student.Password, doc.Password)

		if !password {
			c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Incorrect username or password"})
			return
		}

		token, err := s.jwtService.NewJWT(strconv.FormatUint(uint64(doc.StudentID), 10), doc.Role)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token":      token,
			"expires_in": 60,
		})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "No student found with this id"})
}

func (s *studentHandler) GetStudentByID(c *gin.Context) {

	id := c.Param("id")
	sid, _ := strconv.ParseUint(id, 10, 64)

	if doc, err := s.studentServices.FetchStudentByID(uint(sid)); err == nil {
		m := s.studentMapper.StudentResponseMapper(doc)

		c.JSON(http.StatusOK, gin.H{"status": true, "data": m})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"status": false, "message": "No student found"})

}

func (s *studentHandler) PostTermRegistration(c *gin.Context) {

	id := c.MustGet("id").(string)
	sid, _ := strconv.ParseUint(id, 10, 64)

	var registration dtos.TermRegistrationDTO

	if err := s.validator.Validate(&registration, c); err != nil {
		return
	}

	if err := s.studentServices.CreateTermRegistration(registration, uint(sid)); err == nil {
		c.JSON(http.StatusOK, gin.H{"status": true, "message": "Registration has been sent for approval"})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Error registering for term"})
}

func (s *studentHandler) GetStudentTimetable(c *gin.Context) {

	id := c.MustGet("id").(string)
	sid, _ := strconv.ParseUint(id, 10, 64)

	if doc, err := s.studentServices.FetchStudentTimetable(uint(sid)); err == nil {
		if len(doc) > 0 {
			mappedData := s.studentMapper.StudentTimetableMapper(doc)
			c.JSON(http.StatusOK, gin.H{"status": true, "data": mappedData})
			return
		}

		c.JSON(http.StatusNotFound, gin.H{"status": false, "message": "Not registered for any course"})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Error fetching timetable"})
}

func (s *studentHandler) GetStudentExamSchedule(c *gin.Context) {

	id := c.MustGet("id").(string)
	sid, _ := strconv.ParseUint(id, 10, 64)

	if doc, err := s.studentServices.FetchStudentExamSchedule(uint(sid)); err == nil {
		mappedData := s.studentMapper.StudentExamScheduleMapper(doc)
		c.JSON(http.StatusOK, gin.H{"status": true, "data": mappedData})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"status": false, "message": "No exam schedule found"})
}
