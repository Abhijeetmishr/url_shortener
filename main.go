package main

import (
	handler "gamezop/url_shortner/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", handler.Ping)
	router.Run()
}
