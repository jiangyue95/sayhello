package message

import (
	"sayhello/pkg/db"
)

func Create(msg *Message) error {
	return db.DB.Create(msg).Error
}

func GetAll() ([]Message, error) {
	var messages []Message
	err := db.DB.Find(&messages).Error
	return messages, err
}
