package main

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/api/handlers"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/api/routers"
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
	// secretKey := viper.Get("JWT_SECRET").(string)

	/*
		"""
		Connect to the database
		"""
	*/
	var db *gorm.DB = database.Init(dbUrl)
	defer database.CloseDatabaseConnection(db)

	/*
		"""
		JWT Token Manager
		"""
	*/

	// var jwtService security.TokenManager = security.NewTokenManager(secretKey)

	/*
		"""
		Register All Repositories to db
		"""
	*/

	var (
		studentRepository repositories.StudentRepository = repositories.NewStudentRepository(db)
	)

	/*
		"""
		Register All Mappers
		"""
	*/

	var (
		studentMapper mappers.StudentMapper = mappers.NewStudentMapper()
	)

	/*
		"""
		Register All services
		"""
	*/

	var (
		studentServices services.StudentServices = services.NewStudentServices(studentRepository, studentMapper)
	)

	/*
		"""
		Register All Controllers
		"""
	*/

	var (
		studentHandler handlers.StudentHandler = handlers.NewStudentsHandler(studentServices, studentMapper)
	)

	/*
		"""
		Register the routes to Gin Engine
		"""
	*/

	r := gin.Default()
	api := r.Group("api/v1")

	/*
		"""
		Router without JWT middleware
		"""
	*/

	routers.NoAuthRouter(api)

	/*
		"""
		Router with JWT middleware
		"""
	*/
	auth := r.Group("/api/v1") //, middleware.AuthorizeJWT(jwtService))

	routers.StudentRouter(auth, studentHandler)
	/*
		"""
		Start the server, when you use 'localhost' it will not ask you for the permision again and again "MAC trick"
		"""
	*/

	r.Run("localhost" + port)

}
