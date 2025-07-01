package main

import (
	"log"
	"sayhello/internal/message"
	"sayhello/pkg/db"
)

func AutoMigrate() {
	err := db.DB.AutoMigrate(&message.Message{})
	if err != nil {
		log.Fatalf("Auto migrate Databse Failed: %v", err)
	}

	log.Printf("Auto migrate Database Success")
}
