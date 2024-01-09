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
	productRepository := repository.NewProductRepository(db)
	categoryRepository := repository.NewCategoryRepository(db)

	//service
	userService := service.NewUserService(userRepository)
	productService := service.NewProductService(productRepository, categoryRepository)

	//handler
	userHandler := handler.NewUserHandler(userService)
	productHandler := handler.NewProductHandler(productService)

	router := gin.Default()

	v1 := router.Group("/api/v1")
	v1.POST("/register", userHandler.Register)
	v1.POST("/login", userHandler.Login)

	v1.GET("/category/:category-name/product", productHandler.GetAllProductByNameCategory)

	router.Run()
}
