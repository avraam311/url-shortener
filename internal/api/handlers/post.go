package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/avraam311/url-shortener/internal/models/dto"

	"github.com/wb-go/wbf/ginext"
	"github.com/wb-go/wbf/zlog"
)

func (h *Handler) CreateShortUrl(c *ginext.Context) {
	var req *dto.Request

	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		zlog.Logger.Error().Err(err).Msg("failed to decode request body")
		Fail(c.Writer, http.StatusBadRequest, fmt.Errorf("invalid request body: %s", err.Error()))
		return
	}

	if err := h.validator.Struct(req); err != nil {
		zlog.Logger.Error().Err(err).Msg("failed to validate request body")
		Fail(c.Writer, http.StatusBadRequest, fmt.Errorf("validation error: %s", err.Error()))
		return
	}

	shortURL, err := h.service.CreateShortUrl(c.Request.Context(), req)
	if err != nil {
		zlog.Logger.Error().Err(err).Interface("message", req.FullURL).Msg("failed to create short url")
		Fail(c.Writer, http.StatusInternalServerError, fmt.Errorf("internal server error"))
		return
	}

	Created(c.Writer, shortURL)
}
