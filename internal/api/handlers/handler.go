package handlers

import (
	"context"

	"github.com/avraam311/url-shortener/internal/models/domain"

	"github.com/wb-go/wbf/ginext"

	"github.com/go-playground/validator/v10"
)

type Service interface {
	CreateShortUrl(context.Context, *domain.Url) (string, error)
	GetFullURL(context.Context, string) (string, error)
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

func (h *Handler) GetAnalytics(c *ginext.Context) {}
