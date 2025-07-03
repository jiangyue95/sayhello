package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"sayhello/internal/message"
	"time"
)

func RegisterRoutes() *gin.Engine {
	r := gin.Default()

	// Open CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Accept", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := r.Group("/api")
	{
		api.GET("/messages", message.ListMessagesHandler)
		api.POST("/messages", message.CreateMessageHandler)
	}

	msg := r.Group("/messages")
	{
		msg.POST("", message.CreateMessageHandler)
		msg.GET("", message.ListMessagesHandler)
	}

	r.StaticFile("/", "./static/index.html")

	return r
}
