package main

import (
	"gamezop/url_shortner/internal/config"
	"gamezop/url_shortner/internal/handler"
	"gamezop/url_shortner/internal/repository"
	"gamezop/url_shortner/internal/service"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v\n", err)
	}
	rdb, _ := config.InitRedis()
	pgDb := config.ConnectDatabase()
	h := handler.NewHandler(service.NewSevice(repository.NewRepository(rdb, pgDb)))
	router := gin.Default()
	router.GET("/ping", h.Ping)
	router.POST("/v1/api/create", h.CreateShortUrl)
	router.Run()
}
