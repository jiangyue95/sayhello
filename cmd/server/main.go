package main

import (
	"github.com/spf13/viper"
	"log"

	"sayhello/pkg/db"
)

func main() {
	viper.SetConfigName("config.dev")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	cfg := db.Config{
		User:     viper.GetString("database.user"),
		Password: viper.GetString("database.password"),
		Host:     viper.GetString("database.host"),
		Port:     viper.GetInt("database.port"),
		Name:     viper.GetString("database.name"),
	}

	db.InitMySQL(cfg)
	AutoMigrate()

	r := RegisterRoutes()
	r.Run(":8080")
}
