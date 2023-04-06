package main

import (
	"clean/pkg/config"
	"clean/pkg/db"
	"clean/pkg/di"
	"fmt"
	"log"
)

func main() {

	// env load
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("failed to load env")
	// }

	// configuration
	config, configerr := config.LoadConfig()
	if configerr != nil {
		log.Fatal("failed to load config env", config)
	}

	// connect to database
	gorm, _ := db.ConnectGormDB(config)
	fmt.Println(gorm)
	println("db connected")

	// server
	server, diErr := di.InitializeEvent(config)

	if diErr != nil {

		log.Fatal("cannot start server :", diErr)
	} else {
		server.Start()

	}
}
