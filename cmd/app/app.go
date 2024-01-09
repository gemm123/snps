package main

import (
	"github.com/joho/godotenv"
	"log"
	"synapsis/config"
	"synapsis/internal/handler"
	"synapsis/internal/repository"
	"synapsis/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error load .env")
	}

	db, err := config.NewDbPool()
	if err != nil {
		log.Fatal(err)
	}
	defer config.CloseDB(db)

	//repository
	userRepository := repository.NewUserRepository(db)

	//service
	userService := service.NewUserService(userRepository)

	//handler
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()

	v1 := router.Group("/api/v1")
	v1.POST("/register", userHandler.Register)
	v1.POST("/login", userHandler.Login)

	router.Run()
}
