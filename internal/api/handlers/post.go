package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/avraam311/url-shortener/internal/models/domain"

	"github.com/wb-go/wbf/ginext"
	"github.com/wb-go/wbf/zlog"
)

func (h *Handler) CreateShortUrl(c *ginext.Context) {
	var url *domain.Url

	if err := json.NewDecoder(c.Request.Body).Decode(&url); err != nil {
		zlog.Logger.Error().Err(err).Msg("failed to decode request body")
		Fail(c.Writer, http.StatusBadRequest, fmt.Errorf("invalid request body: %s", err.Error()))
		return
	}

	if err := h.validator.Struct(url); err != nil {
		zlog.Logger.Error().Err(err).Msg("failed to validate request body")
		Fail(c.Writer, http.StatusBadRequest, fmt.Errorf("validation error: %s", err.Error()))
		return
	}

	shortURL, err := h.service.CreateShortUrl(c.Request.Context(), url)
	if err != nil {
		zlog.Logger.Error().Err(err).Interface("message", url.Url).Msg("failed to create short url")
		Fail(c.Writer, http.StatusInternalServerError, fmt.Errorf("internal server error"))
		return
	}

	Created(c.Writer, shortURL)
}
