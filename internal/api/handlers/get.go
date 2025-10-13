package handlers

import (
	"fmt"
	"net/http"

	"github.com/wb-go/wbf/ginext"
	"github.com/wb-go/wbf/zlog"
)

func (h *Handler) GoToShortUrl(c *ginext.Context) {
	shortURL := c.Param("short_url")

	if shortURL == "" {
		zlog.Logger.Error().Msg("empty short url")
		Fail(c.Writer, http.StatusBadRequest, fmt.Errorf("empty short url"))
		return
	}

	fullURL, err := h.service.GetFullURL(c.Request.Context(), shortURL)
	if err != nil {
		zlog.Logger.Error().Err(err).Msg("failed to get full url")
		Fail(c.Writer, http.StatusInternalServerError, fmt.Errorf("internal server error"))
		return
	}

	c.Redirect(http.StatusFound, fullURL)
}

func (h *Handler) GetAnalytics(c *ginext.Context) {
}
