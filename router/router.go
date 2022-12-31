package router

import (
	"net/http"
	"rabbitmq-producer-site-check/constants"
	"rabbitmq-producer-site-check/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//SetupRoutes : gin router
func SetupRoutes() {
	httpRouter := gin.Default()

	//handling CORS
	httpRouter.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	httpRouter.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Golang RabbitMQ API 📺 Up and Running"})
	})

	httpRouter.POST("task1", controller.Task1Controller)
	httpRouter.POST("task2", controller.Task2Controller)

	httpRouter.Run(constants.SERVERPORT)

}
