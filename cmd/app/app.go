package main

import (
	"github.com/joho/godotenv"
	"log"
	"synapsis/config"
	"synapsis/internal/handler"
	"synapsis/internal/repository"
	"synapsis/internal/service"
	"synapsis/middleware"

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
	cartRepository := repository.NewCartRepository(db)
	orderRepository := repository.NewOrderRepository(db)

	//service
	userService := service.NewUserService(userRepository)
	productService := service.NewProductService(productRepository, categoryRepository)
	cartService := service.NewCartService(cartRepository, productRepository)
	orderService := service.NewOrderService(orderRepository, cartRepository, productRepository)

	//handler
	userHandler := handler.NewUserHandler(userService)
	productHandler := handler.NewProductHandler(productService)
	cartHandler := handler.NewCartHandler(cartService)
	orderHandler := handler.NewOrderHandler(orderService, cartService)

	router := gin.Default()

	v1 := router.Group("/api/v1")
	v1.POST("/register", userHandler.Register)
	v1.POST("/login", userHandler.Login)

	v1.GET("/category/:category-name/product", middleware.Auth(), productHandler.GetAllProductByNameCategory)

	v1.POST("/cart", middleware.Auth(), cartHandler.AddProductToCart)
	v1.GET("/cart", middleware.Auth(), cartHandler.GetCartProduct)
	v1.DELETE("/cart/delete/:id-product", middleware.Auth(), cartHandler.DeleteCartProduct)

	v1.POST("/order", middleware.Auth(), orderHandler.AddOrderProduct)
	v1.GET("/order", middleware.Auth(), orderHandler.GetOrderProduct)

	router.Run()
}
