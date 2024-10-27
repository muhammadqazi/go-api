package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/muhammadqazi/campus-hq-api/src/internal/api/handlers"
	middleware "github.com/muhammadqazi/campus-hq-api/src/internal/api/middlewares"
	"github.com/muhammadqazi/campus-hq-api/src/internal/api/routers"
	"github.com/muhammadqazi/campus-hq-api/src/internal/common/security"
	"github.com/muhammadqazi/campus-hq-api/src/internal/common/validation"
	core "github.com/muhammadqazi/campus-hq-api/src/internal/config"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/domain/services"
	database "github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/mappers"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/repositories"
	"gorm.io/gorm"
)

func main() {
	/*
		"""
		Load the environment variables from .env file
		"""
	*/

	conf, err := core.LoadConfig()
	if err != nil {
		errfmt := "Error loading config: %s"
		panic(fmt.Errorf(errfmt, err))
	}

	/*
		"""
		Connect to the database
		"""
	*/
	var db *gorm.DB = database.Init(conf.DBUrl)
	defer database.CloseDatabaseConnection(db)

	/*
		"""
		JWT Token Manager & validator
		"""
	*/

	var jwtService security.TokenManager = security.NewTokenManager(conf)
	var validator validation.Validator = validation.NewStructValidator()

	/*
		"""
		Register All Mappers
		"""
	*/

	var (
		studentMapper    mappers.StudentMapper     = mappers.NewStudentMapper()
		accountsMapper   mappers.AccountsMapper    = mappers.NewAccountingMapper()
		curriculumMapper mappers.CurriculumMapper  = mappers.NewCurriculumMapper()
		instructorMapper mappers.InstructorsMapper = mappers.NewInstructorsMapper()
		courseMapper     mappers.CourseMapper      = mappers.NewCourseMapper()
		examMapper       mappers.ExamMapper        = mappers.NewExamMapper()
		facultyMapper    mappers.FacultyMapper     = mappers.NewFacultyMapper()
		departmentMapper mappers.DepartmentMapper  = mappers.NewDepartmentMapper()
		buildingMapper   mappers.BuildingMapper    = mappers.NewBuildingMapper()
		roomMapper       mappers.RoomMapper        = mappers.NewRoomMapper()
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
		facultyRepository    repositories.FacultyRepository     = repositories.NewFacultyRepository(db, facultyMapper)
		departmentRepository repositories.DepartmentRepository  = repositories.NewDepartmentRepository(db, departmentMapper)
		buildingRepository   repositories.BuildingRepository    = repositories.NewBuildingRepository(db, buildingMapper)
		roomRepository       repositories.RoomRepository        = repositories.NewRoomRepository(db, roomMapper)
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
		facultyServices    services.FacultyServices     = services.NewFacultyServices(facultyRepository, facultyMapper)
		departmentServices services.DepartmentServices  = services.NewDepartmentServices(departmentRepository, departmentMapper)
		buildingServices   services.BuildingServices    = services.NewBuildingServices(buildingRepository, buildingMapper)
		roomServices       services.RoomServices        = services.NewRoomServices(roomRepository, roomMapper)
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
		facultyHandler    handlers.FacultyHandler     = handlers.NewFacultyHandler(facultyServices, facultyMapper, validator)
		departmentHandler handlers.DepartmentHandler  = handlers.NewDepartmentHandler(departmentServices, departmentMapper, validator)
		buildingHandler   handlers.BuildingHandler    = handlers.NewBuildingHandler(buildingServices, buildingMapper, validator)
		roomHandler       handlers.RoomHandler        = handlers.NewRoomHandler(roomServices, roomMapper, validator)
	)

	/*
		"""
		Register the routes to Gin Engine
		"""
	*/

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Use(middleware.CORSMiddleware())

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
	routers.FacultyRouter(auth, facultyHandler)
	routers.DepartmentRouter(auth, departmentHandler)
	routers.BuildingRouter(auth, buildingHandler)
	routers.RoomRouter(auth, roomHandler)

	/*
		"""
		Start the server, when you use 'localhost' it will not ask you for the permission again and again "MAC trick"
		"""
	*/

	r.Run("0.0.0.0" + conf.Port)

}
