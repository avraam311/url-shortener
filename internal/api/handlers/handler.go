package handlers

import (
	"context"

	"github.com/avraam311/url-shortener/internal/models/db"
	"github.com/avraam311/url-shortener/internal/models/dto"

	"github.com/go-playground/validator/v10"
)

type Service interface {
	CreateShortUrl(context.Context, *dto.Request) (string, error)
	GetFullURL(context.Context, string) (string, error)
	GetAnalytics(context.Context, string) ([]*db.Analytics, error)
}

type Handler struct {
	service   Service
	validator *validator.Validate
}

func NewHandler(service Service, validator *validator.Validate) *Handler {
	return &Handler{
		service:   service,
		validator: validator,
	}
}
