package message

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateMessageHandler(c *gin.Context) {
	var msg Message
	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := CreateMessage(&msg); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "save failed"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": msg})
}

func ListMessagesHandler(c *gin.Context) {
	msgs, err := ListMessages()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "query failed"})
		return
	}

	var formattedMsgs []gin.H
	for _, msg := range msgs {
		formattedMsgs = append(formattedMsgs, gin.H{
			"id":         msg.ID,
			"name":       msg.Name,
			"body":       msg.Body,
			"created_at": formatTime(msg.CreatedAt),
		})
	}
	c.JSON(http.StatusOK, formattedMsgs)
}
