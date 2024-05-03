package internal

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mitsuha/stork/internal/services/albums"
	"github.com/mitsuha/stork/internal/services/overview"
	"github.com/mitsuha/stork/internal/services/reverseProxy"
	"github.com/mitsuha/stork/pkg/authentication"
	"time"
)

func Run() error {
	engine := gin.Default()

	c := cors.New(cors.Config{
		AllowAllOrigins: true,
		//AllowMethods:     []string{"PUT", "PATCH"},
		//AllowHeaders:     []string{"Origin"},
		//ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})

	r := engine.Group("/api", c, authentication.Auth)
	{
		service := overview.New()

		r.GET("/data", service.Data)
		r.GET("/overview", service.Overview)
	}
	{
		service := albums.New()

		router := r.Group("/albums")
		router.GET("/:id", service.Show)
		router.GET("/:id/songs", service.Songs)
	}

	engine.NoRoute(reverseProxy.New())

	return engine.Run()
}
