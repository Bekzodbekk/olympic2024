package http

import (
	"api-gateway/internal/http/handler"
	service "api-gateway/internal/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @tite Api-gateway service
// @version 1.0
// @description Api-gateway service
// @host localhost:9000
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authourization
func NewGin(service *service.ServiceRepositoryClient) *gin.Engine {

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	handler := handler.NewHandler(service)

	// Authentication routes
	r.POST("/auth/register", handler.RegisterUser)
	r.POST("/auth/login", handler.LoginUser)
	r.POST("/auth/refresh", handler.RefreshToken)

	// User routes
	r.PUT("/users/:id", handler.UpdateUser)
	r.GET("/users/:id", handler.GetUserById)
	r.GET("/users", handler.GetUsers)
	r.GET("/users/filter", handler.GetUserByFilter)

	// Athlete routes
	r.POST("/athletes", handler.CreateAthlete)
	r.GET("/athletes/:id", handler.GetAthlete)
	r.GET("/athletes", handler.ListOfAthlete)
	r.PUT("/athletes/:id", handler.UpdateAthlete)
	r.DELETE("/athletes/:id", handler.DeleteAthlete)

	// Event routes
	r.POST("/events", handler.CreateEvent)
	r.GET("/events/:id", handler.GetEvent)
	r.GET("/events", handler.ListOfEvent)
	r.PUT("/events/:id", handler.UpdateEvent)
	r.DELETE("/events/:id", handler.DeleteEvent)

	// Country routes
	r.POST("/countries", handler.CreateCountry)
	r.GET("/countries/:id", handler.GetCountry)
	r.GET("/countries", handler.ListOfCountry)
	r.PUT("/countries/:id", handler.UpdateCountry)
	r.DELETE("/countries/:id", handler.DeleteCountry)



	return r
}