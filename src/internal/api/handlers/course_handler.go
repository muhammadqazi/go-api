package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadqazi/campus-hq-api/src/internal/common/validation"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/domain/services"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/mappers"
	"net/http"
	"strconv"
)

/*
	"""
	CourseHandler can provide the following services.
	"""
*/

type CourseHandler interface {
	PostCourse(c *gin.Context)
	GetCourseByCourseCode(c *gin.Context)
	PatchCourseByCourseCode(c *gin.Context)
	DeleteCourseByCourseCode(c *gin.Context)
	PatchCourseInstructorByCourseId(c *gin.Context)
}

type courseHandler struct {
	validator      validation.Validator
	courseMapper   mappers.CourseMapper
	courseServices services.CourseServices
}

/*
	"""
	This will create a new instance of the CourseHandler, we will use this as a constructor
	"""
*/

func NewCourseHandler(service services.CourseServices, mapper mappers.CourseMapper, v validation.Validator) CourseHandler {
	return &courseHandler{
		courseMapper:   mapper,
		courseServices: service,
		validator:      v,
	}
}

func (s *courseHandler) PostCourse(c *gin.Context) {
	var course dtos.CourseCreateDTO

	if err := s.validator.Validate(&course, c); err != nil {
		return
	}

	if err := s.courseServices.CreateCourse(course); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Course created successfully"})
}

func (s *courseHandler) GetCourseByCourseCode(c *gin.Context) {

	code := c.Param("code")

	if course, err := s.courseServices.FetchCourseByCourseCode(code); err == nil {
		mapped := s.courseMapper.CourseFetchByCodeMapper(course)
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "data": mapped})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Error fetching course"})
}

func (s *courseHandler) PatchCourseByCourseCode(c *gin.Context) {
	code := c.Param("code")
	var course dtos.CourseUpdateDTO

	if err := s.validator.Validate(&course, c); err != nil {
		return
	}

	if err := s.courseServices.ModifyCourseByCourseCode(code, course); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Course updated successfully"})
}

func (s *courseHandler) DeleteCourseByCourseCode(c *gin.Context) {
	code := c.Param("code")

	if err := s.courseServices.RemoveCourseByCourseCode(code); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Course deleted successfully"})
}

func (s *courseHandler) PatchCourseInstructorByCourseId(c *gin.Context) {
	id := c.Param("id")
	uintCourseId, _ := strconv.ParseUint(id, 10, 64)

	var course dtos.CourseInstructorUpdateDTO
	course.CourseId = uint(uintCourseId)

	if err := s.validator.Validate(&course, c); err != nil {
		return
	}

	if err := s.courseServices.ModifyCourseInstructorByCourseId(course); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Course instructor updated successfully"})
}
