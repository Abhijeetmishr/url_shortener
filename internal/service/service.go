package service

import (
	"gamezop/url_shortner/internal/model"
	repository "gamezop/url_shortner/internal/repository"

	"github.com/gin-gonic/gin"
)

type Service struct {
	r *repository.Repository
}

func NewSevice(repo *repository.Repository) *Service {
	return &Service{r: repo}
}

func (s *Service) CreateShortUrl(ctx *gin.Context, longUrl model.LongURL) (string, error) {
	// check if exists in redis
	// if not check for postgres db -> if exists return else
	// create a short url for the corresponding longUrl -> insert in db & update redis cache
	//
	return "", nil
}
