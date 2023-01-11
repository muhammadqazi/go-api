package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/common/security"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/domain/services"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/mappers"
)

/*
	"""
	StudentHandler can provide the following services.
	"""
*/

type StudentHandler interface {
	CreateStudent(c *gin.Context)
	GetStudentByEmail(c *gin.Context)
	GetStudentByStudentID(c *gin.Context)
	StudentSignIn(c *gin.Context)
}

type studentHandler struct {
	studentMapper   mappers.StudentMapper
	studentServices services.StudentServices
	jwtService      security.TokenManager
}

/*
	"""
	This will creates a new instance of the StudentHandler, we will use this as a constructor
	"""
*/

func NewStudentsHandler(service services.StudentServices, mapper mappers.StudentMapper, jwtService security.TokenManager) StudentHandler {
	return &studentHandler{
		studentMapper:   mapper,
		studentServices: service,
		jwtService:      jwtService,
	}
}

func (s *studentHandler) CreateStudent(c *gin.Context) {

	var student dtos.StudentCreateDTO

	/*
		"""
		BindJSON will bind the request body to the struct
		"""
	*/

	if err := c.BindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
		return
	}

	/*

		"""
		We will check if the student already exists in the database
		"""
	*/

	if _, err := s.studentServices.GetStudentByEmail(student.Email); err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Student already exists"})
		return
	}

	recent_sid, err := s.studentServices.GetLastStudentID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Error generating student number"})
		return
	}

	semester := getCurrentSemester()
	sid := recent_sid + 1

	if sid, err := s.studentServices.CreateStudent(student, sid, semester); err == nil {
		c.JSON(http.StatusCreated, gin.H{"status": true, "message": "Student created successfully", "student_id": sid})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Error creating student"})

}

func (s *studentHandler) GetStudentByEmail(c *gin.Context) {

}

func (s *studentHandler) StudentSignIn(c *gin.Context) {

	var student dtos.StudentSignInDTO

	/*
		"""
		BindJSON will bind the request body to the struct
		"""
	*/

	if err := c.BindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
		return
	}

	if doc, err := s.studentServices.GetStudentByStudentID(student.StudentID); err == nil {

		password := security.CheckPasswordHash(student.Password, doc.Password)

		if !password {
			c.JSON(http.StatusBadRequest, gin.H{"status": false, "messsage": "Incorrect username or password"})
			return
		}

		token, err := s.jwtService.NewJWT(strconv.FormatUint(uint64(doc.StudentID), 10))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":     true,
			"token":      token,
			"expires_in": 60,
		})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "No student found with this id"})
}

func (s *studentHandler) GetStudentByStudentID(c *gin.Context) {

	id := c.Param("id")
	sid, _ := strconv.ParseUint(id, 10, 64)

	if doc, err := s.studentServices.GetStudentByStudentID(uint(sid)); err == nil {
		m := s.studentMapper.StudentResponseMapper(doc)

		c.JSON(http.StatusOK, gin.H{"status": true, "data": m})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"status": false, "message": "No student found"})

}

func getCurrentSemester() string {
	_, month, _ := time.Now().Date()
	var semester string
	switch {
	case month >= time.September && month <= time.January:
		semester = "Fall"
	case month >= time.February && month <= time.June:
		semester = "Spring"
	default:
		semester = "Summer"
	}

	return semester
}
