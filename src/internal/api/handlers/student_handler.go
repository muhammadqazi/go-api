package handlers

import (
	"errors"
	"github.com/muhammadqazi/campus-hq-api/src/internal/common/utils"
	"github.com/muhammadqazi/campus-hq-api/src/internal/common/validation"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/muhammadqazi/campus-hq-api/src/internal/common/security"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/domain/services"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/mappers"
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
	GetStudentCoursesAttendance(c *gin.Context)
	PatchResetPassword(c *gin.Context)
	PostForgotPasswordRequest(c *gin.Context)
	PutForgotPasswordCode(c *gin.Context)
	PatchNewPassword(c *gin.Context)
	TestApi(c *gin.Context)
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

func (s *studentHandler) TestApi(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": true, "message": "API is working"})
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
			"role":       doc.Role,
			"student_id": doc.StudentID,
			"email":      doc.Email,
			"first_name": doc.FirstName,
			"last_name":  doc.Surname,
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

	isExist, err := s.studentServices.FetchIsEnrollmentExists(uint(sid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}

	if isExist {
		isEnrolled, err := s.studentServices.FetchStudentEnrollmentStatus(uint(sid))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
			return
		}

		if isEnrolled {
			c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "You have already registered for this term"})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "You have already sent a registration request"})
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

func (s *studentHandler) GetStudentCoursesAttendance(c *gin.Context) {

	id := c.MustGet("id").(string)
	sid, _ := strconv.ParseUint(id, 10, 64)

	if doc, err := s.studentServices.FetchStudentCoursesAttendance(uint(sid)); err == nil {

		if len(doc) > 0 {
			mappedData := s.studentMapper.StudentCourseAttendanceFetchMapper(doc)
			c.JSON(http.StatusOK, gin.H{"status": true, "data": mappedData})
			return
		}

		c.JSON(http.StatusNotFound, gin.H{"status": false, "message": "No attendance found"})
		return

	}

	c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Error fetching attendance"})
}

func (s *studentHandler) PatchResetPassword(c *gin.Context) {
	var password dtos.ResetPasswordDTO

	if err := s.validator.Validate(&password, c); err != nil {
		return
	}

	id := c.MustGet("id").(string)
	sid, _ := strconv.ParseUint(id, 10, 64)

	student, err := s.studentServices.FetchStudentByID(uint(sid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}

	if match := security.CheckPasswordHash(password.CurrentPassword, student.Password); !match {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Invalid old password"})
		return
	}

	if hashedPassword, err := security.HashPassword(password.NewPassword); err == nil {
		if err := s.studentServices.ModifyStudentPassword(uint(sid), hashedPassword); err == nil {
			c.JSON(http.StatusOK, gin.H{"status": true, "message": "Password updated successfully"})
			return
		}
	}

	c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Error updating password"})

}

func (s *studentHandler) PostForgotPasswordRequest(c *gin.Context) {
	var request dtos.ForgotPasswordRequestDTO

	if err := s.validator.Validate(&request, c); err != nil {
		return
	}

	_, err := s.studentServices.FetchStudentByID(request.StudentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}

	if err := s.studentServices.CreateForgotPasswordRequest(request); err == nil {
		token, _ := s.jwtService.NewPasswordResetJWT(strconv.FormatUint(uint64(request.StudentID), 10))

		c.JSON(http.StatusOK, gin.H{"status": true, "token": token})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Error sending request"})

}

func (s *studentHandler) PutForgotPasswordCode(c *gin.Context) {

	var code dtos.ForgotPasswordVerifyDTO

	if err := s.validator.Validate(&code, c); err != nil {
		return
	}

	id := c.MustGet("id").(string)
	sid, _ := strconv.ParseUint(id, 10, 64)

	if doc, err := s.studentServices.VerifyForgotPasswordCode(uint(sid)); err == nil {

		if doc.ExpiresAt.Before(time.Now()) {
			c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Code has been expired"})
			return
		}

		if doc.IsVerified {
			c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Code has already been used"})
			return
		}

		if doc.ResetCode != code.Code {
			c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Invalid code"})
			return
		}

		if err := s.studentServices.ModifyForgotPasswordFlag(uint(sid)); err == nil {
			c.JSON(http.StatusOK, gin.H{"status": true, "message": "Code verified successfully"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Error verifying code"})
		return

	}

	c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Invalid code"})

}

func (s *studentHandler) PatchNewPassword(c *gin.Context) {

	var password dtos.NewPasswordDTO

	if err := s.validator.Validate(&password, c); err != nil {
		return
	}

	id := c.MustGet("id").(string)
	sid, _ := strconv.ParseUint(id, 10, 64)

	if hashedPassword, err := security.HashPassword(password.Password); err == nil {
		if doc, err := s.studentServices.VerifyForgotPasswordCode(uint(sid)); err == nil {
			if doc.IsVerified {
				if err := s.studentServices.ModifyStudentPassword(uint(sid), hashedPassword); err == nil {
					if err := s.studentServices.RemoveForgotPasswordCode(uint(sid)); err == nil {
						c.JSON(http.StatusOK, gin.H{"status": true, "message": "Password updated successfully"})
						return
					}
				}
			}
			c.JSON(http.StatusUnauthorized, gin.H{"status": false, "message": "Unauthorized"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "No password reset request found"})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Error updating password"})
}
