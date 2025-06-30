package message

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Name string
	body string
}
