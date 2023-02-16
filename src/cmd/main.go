package main

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/api/handlers"
	middleware "github.com/muhammadqazi/SIS-Backend-Go/src/internal/api/middlewares"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/api/routers"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/common/security"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/common/validation"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/domain/services"
	database "github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/mappers"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/repositories"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func main() {

	/*
		"""
		Load the environment variables from .env file
		"""
	*/

	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)
	secretKey := viper.Get("JWT_SECRET").(string)

	/*
		"""
		Connect to the database
		"""
	*/
	var db *gorm.DB = database.Init(dbUrl)
	defer database.CloseDatabaseConnection(db)

	/*
		"""
		JWT Token Manager & validator
		"""
	*/

	var jwtService security.TokenManager = security.NewTokenManager(secretKey)
	var validator validation.Validator = validation.NewStructValidator()

	/*
		"""
		Register All Mapperss
		"""
	*/

	var (
		studentMapper    mappers.StudentMapper     = mappers.NewStudentMapper()
		accountsMapper   mappers.AccountsMapper    = mappers.NewAccountingMapper()
		curriculumMapper mappers.CurriculumMapper  = mappers.NewCurriculumMapper()
		instructorMapper mappers.InstructorsMapper = mappers.NewInstructorsMapper()
		courseMapper     mappers.CourseMapper      = mappers.NewCourseMapper()
		examMapper       mappers.ExamMapper        = mappers.NewExamMapper()
	)

	/*
		"""
		Register All Repositories to db
		"""
	*/

	var (
		studentRepository    repositories.StudentRepository     = repositories.NewStudentRepository(db, studentMapper)
		accountsRepository   repositories.AccountingRepository  = repositories.NewAccountingRepository(db)
		curriculumRepository repositories.CurriculumRepository  = repositories.NewCurriculumRepository(db, curriculumMapper)
		instructorRepository repositories.InstructorsRepository = repositories.NewInstructorsRepository(db, instructorMapper)
		courseRepository     repositories.CourseRepository      = repositories.NewCourseRepository(db, courseMapper)
		examRepository       repositories.ExamRepository        = repositories.NewExamRepository(db, examMapper)
	)

	/*
		"""
		Register All services
		"""
	*/

	var (
		studentServices    services.StudentServices     = services.NewStudentServices(studentRepository, studentMapper)
		accountServices    services.AccountingServices  = services.NewAccountingServices(accountsRepository, accountsMapper)
		curriculumServices services.CurriculumServices  = services.NewCurriculumServices(curriculumRepository, curriculumMapper)
		instructorServices services.InstructorsServices = services.NewInstructorsServices(instructorRepository, instructorMapper)
		courseServices     services.CourseServices      = services.NewCourseServices(courseRepository, courseMapper)
		examServices       services.ExamServices        = services.NewExamServices(examRepository, examMapper)
	)

	/*
		"""
		Register All Controllers
		"""
	*/

	var (
		studentHandler    handlers.StudentHandler     = handlers.NewStudentsHandler(studentServices, accountServices, studentMapper, jwtService, validator)
		accountsHandler   handlers.AccountingHandler  = handlers.NewAccountingHandler(accountServices, accountsMapper, validator)
		curriculumHandler handlers.CurriculumHandler  = handlers.NewCurriculumHandler(curriculumServices, curriculumMapper, validator)
		instructorHandler handlers.InstructorsHandler = handlers.NewInstructorsHandler(instructorServices, instructorMapper, jwtService, validator)
		courseHandler     handlers.CourseHandler      = handlers.NewCourseHandler(courseServices, courseMapper, validator)
		examHandler       handlers.ExamHandler        = handlers.NewExamHandler(examServices, examMapper, validator)
	)

	/*
		"""
		Register the routes to Gin Engine
		"""
	*/

	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	/*
		"""
		Router with JWT middleware
		"""
	*/
	auth := r.Group("/api/v1", middleware.AuthorizeJWT(jwtService))

	routers.StudentRouter(auth, studentHandler)
	routers.AccountingRouter(auth, accountsHandler)
	routers.CurriculumRouter(auth, curriculumHandler)
	routers.InstructorsRouter(auth, instructorHandler)
	routers.CourseRouter(auth, courseHandler)
	routers.ExamRouter(auth, examHandler)

	/*
		"""
		Start the server, when you use 'localhost' it will not ask you for the permision again and again "MAC trick"
		"""
	*/

	r.Run("localhost" + port)

}
