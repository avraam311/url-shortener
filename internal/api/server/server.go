package server

import (
	"net/http"

	"github.com/wb-go/wbf/ginext"

	"github.com/avraam311/url-shortener/internal/api/handlers"
	"github.com/avraam311/url-shortener/internal/middlewares"
)

func NewRouter(handler *handlers.Handler, ginMode string) *ginext.Engine {
	e := ginext.New(ginMode)

	e.Use(middlewares.CORSMiddleware())
	e.Use(ginext.Logger())
	e.Use(ginext.Recovery())

	api := e.Group("/api/url-shortener")
	{
		api.POST("/shorten", handler.CreateShortUrl)
		api.GET("/", handler.GoToUrl)
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
