package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/common/security"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/common/validation"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/domain/services"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/mappers"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

/*
	"""
	InstructorsHandler can provide the following services.
	"""
*/

type InstructorsHandler interface {
	PostInstructors(c *gin.Context)
	PostSignIn(c *gin.Context)
	GetTermEnrollmentRequests(c *gin.Context)
	PatchTermEnrollmentRequests(c *gin.Context)
	PostInstructorCourseEnrollment(c *gin.Context)
	GetInstructorCourseEnrollment(c *gin.Context)
	PatchStudentAttendance(c *gin.Context)
}

type instructorsHandler struct {
	validator           validation.Validator
	instructorsMapper   mappers.InstructorsMapper
	instructorsServices services.InstructorsServices
	jwtService          security.TokenManager
}

/*
	"""
	This will create a new instance of the InstructorsHandler, we will use this as a constructor
	"""
*/

func NewInstructorsHandler(service services.InstructorsServices, mapper mappers.InstructorsMapper, jwtService security.TokenManager, v validation.Validator) InstructorsHandler {
	return &instructorsHandler{
		instructorsMapper:   mapper,
		instructorsServices: service,
		jwtService:          jwtService,
		validator:           v,
	}
}

func (s *instructorsHandler) PostInstructors(c *gin.Context) {
	var instructor dtos.InstructorCreateDTO

	if err := s.validator.Validate(&instructor, c); err != nil {
		return
	}

	_, err := s.instructorsServices.FetchInstructorByEmail(instructor.Email)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Email already associated with another instructor"})
		return
	}

	_, err = s.instructorsServices.FetchInstructorByPhone(instructor.PhoneNumber)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Phone number already associated with another instructor"})
		return
	}

	if err := s.instructorsServices.CreateInstructors(instructor); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "There was an error performing this action"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": true, "message": "Instructor created successfully"})
}

func (s *instructorsHandler) PostSignIn(c *gin.Context) {

	var instructor dtos.InstructorSignInDTO

	if err := s.validator.Validate(&instructor, c); err != nil {
		return
	}

	if doc, err := s.instructorsServices.FetchInstructorByEmail(instructor.Email); err == nil {
		password := security.CheckPasswordHash(instructor.Password, doc.Password)

		if !password {
			c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Incorrect email or password"})
			return
		}

		token, err := s.jwtService.NewJWT(strconv.FormatUint(uint64(doc.InstructorID), 10), doc.Role)
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

	c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "No instructor found with this email"})
}

func (s *instructorsHandler) GetTermEnrollmentRequests(c *gin.Context) {

	id := c.MustGet("id")
	instructorID, _ := strconv.ParseUint(id.(string), 10, 64)

	if doc, err := s.instructorsServices.FetchTermEnrollmentRequests(uint(instructorID)); err == nil {

		if len(doc) > 0 {
			mapped := s.instructorsMapper.InstructorTermRequestsMapper(doc)

			c.JSON(http.StatusOK, gin.H{"status": true, "data": mapped})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "No enrollment requests found"})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "There was an error performing this action"})
}

func (s *instructorsHandler) PatchTermEnrollmentRequests(c *gin.Context) {

	var request dtos.InstructorApproveEnrollmentRequestDTO

	id := c.Param("request-id")
	requestID, _ := strconv.ParseUint(id, 10, 64)

	request.RequestID = uint(requestID)

	if err := s.validator.Validate(&request, c); err != nil {
		return
	}
	if err := s.instructorsServices.ModifyTermEnrollmentRequests(request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "There was an error performing this action"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Enrollment request approved successfully"})
}

func (s *instructorsHandler) PostInstructorCourseEnrollment(c *gin.Context) {

	var request dtos.InstructorCourseEnrollmentDTO

	if err := s.validator.Validate(&request, c); err != nil {
		return
	}

	if err := s.instructorsServices.CreateInstructorCourseEnrollment(request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "There was an error performing this action"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": true, "message": "Instructor course enrollment created successfully"})
}

func (s *instructorsHandler) GetInstructorCourseEnrollment(c *gin.Context) {
	id := c.Param("id")
	instructorID, _ := strconv.ParseUint(id, 10, 64)

	if doc, err := s.instructorsServices.FetchInstructorCourseEnrollment(uint(instructorID)); err == nil {

		mappedData := s.instructorsMapper.InstructorFetchCoursesMapper(doc)
		c.JSON(http.StatusOK, gin.H{"status": true, "data": mappedData})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "There was an error performing this action"})
}

func (s *instructorsHandler) PatchStudentAttendance(c *gin.Context) {

	var attendance dtos.StudentAttendancePatchDTO

	if err := s.validator.Validate(&attendance, c); err != nil {
		return
	}

	if err := s.instructorsServices.ModifyStudentAttendance(attendance); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "There was an error performing this action"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Student attendance updated successfully"})
}
