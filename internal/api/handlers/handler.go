package handlers

import (
	"github.com/wb-go/wbf/ginext"
)

type Service interface{}

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) CreateShortUrl(c *ginext.Context) {}

func (h *Handler) GoToUrl(c *ginext.Context) {}

func (h *Handler) GetAnalytics(c *ginext.Context) {}
