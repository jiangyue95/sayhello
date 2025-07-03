package message

import (
	"gorm.io/gorm"
	"time"
)

type Message struct {
	gorm.Model
	Name      string    `json:"name"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

func formatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
