package handler

import (
	model "gamezop/url_shortner/internal/model"
	"gamezop/url_shortner/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	s *service.Service
}

func NewHandler(serv *service.Service) *Handler {
	return &Handler{s: serv}
}

func (h *Handler) Ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}

func (h *Handler) CreateShortUrl(ctx *gin.Context) {
	var req model.LongURL
	if err := ctx.ShouldBind(&req); err != nil {
		log.Println("Invalid JSON input validation failed")
		ctx.JSON(404, gin.H{"error": true, "message": "Input JSON validation failed"})
		return
	}
	var s, err = h.s.CreateShortUrl(ctx, req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": true, "message": "Failed to generate short-url"})
		return
	}
	ctx.JSON(200, gin.H{"error": false, "data": &s})
}
