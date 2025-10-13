package server

import (
	"context"
	"net/http"

	"github.com/wb-go/wbf/ginext"

	"github.com/avraam311/url-shortener/internal/api/handlers"
	"github.com/avraam311/url-shortener/internal/middlewares"
	"github.com/avraam311/url-shortener/internal/models/db"
)

type RepositoryAnalytics interface {
	SaveAnalytics(context.Context, *db.Analytics) error
}

func NewRouter(handler *handlers.Handler, ginMode string, repoAnalytics RepositoryAnalytics) *ginext.Engine {
	e := ginext.New(ginMode)

	e.Use(middlewares.CORSMiddleware())
	e.Use(ginext.Logger())
	e.Use(ginext.Recovery())

	api := e.Group("/api/url-shortener")
	{
		api.POST("/shorten", handler.CreateShortUrl)
		api.GET("/", middlewares.AnalyticsMiddleware(repoAnalytics), handler.GoToShortUrl)
		api.GET("/analytics/:short_url", handler.GetAnalytics)
	}

	return e
}

func NewServer(addr string, router *ginext.Engine) *http.Server {
	return &http.Server{
		Addr:    addr,
		Handler: router,
	}
}
