package main

import (
	"github.com/gin-gonic/gin"
	"sayhello/internal/message"
)

func RegisterRoutes() *gin.Engine {
	r := gin.Default()

	msg := r.Group("/messages")
	{
		msg.POST("", message.CreateMessageHandler)
		msg.GET("", message.ListMessagesHandler)
	}

	r.StaticFile("/", "./static/index.html")

	return r
}
