package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/swaggo/swag/example/celler/httputil"
	_ "github.com/swaggo/swag/example/celler/model"
	_ "realEstate/cmd/server/docs"
	"realEstate/internal/middleware"
	logg "realEstate/pkg/log"
)

var router *gin.Engine

func StartGin() {
	router = gin.Default()
	//corsConfig := cors.DefaultConfig()
	//corsConfig.AllowCredentials = true
	//corsConfig.AllowOrigins = []string{"https:/localhost:8080"}
	// To be able to send tokens to the server.
	//corsConfig.AllowCredentials = true
	// OPTIONS method for ReactJS
	//corsConfig.AddAllowMethods("OPTIONS")
	//router.Use(cors.New(corsConfig))
	InitializeRoutes()
	if err := router.Run("localhost:8080"); err != nil {
		logg.Fatal("Server not run!")
	}
	logg.Info("Server is running")
}

// func initialize routes for server
func InitializeRoutes() {
	router.GET("/", middleware.ShowIndexPage)
	//user service
	router.GET("/users", middleware.GetAllUser)
	router.POST("/auth/register", middleware.CreateUser)
	router.GET("/auth/login", middleware.Login)
	router.GET("/auth/logout", middleware.Logout)
	router.GET("/getusers", middleware.GetUser)
	router.PUT("/users", middleware.UpdateUser)
	router.DELETE("/users", middleware.DeleteUser)
	router.GET(" ", middleware.NotFound)
	//advertisment service
	router.GET("/advertisment", middleware.GetAllAdvertisment)
	router.POST("/advertisment", middleware.CreateAdvertisment)
	router.GET("/getadvertisment", middleware.GetAdvertisment)
	router.PUT("/advertisment", middleware.UpdateAdvertisment)
	router.DELETE("/advertisment", middleware.DeleteAdvertisment)
	router.POST("/upload", middleware.UploadFiles)
	//TODO настроить корректную работу свагера
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/auth/validate",middleware.IsValidUserJWT)
	logg.Info("Routes initialized")
	// content service
	router.GET("/content", middleware.GetContent)
}
