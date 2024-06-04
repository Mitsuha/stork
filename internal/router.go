package internal

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/mitsuha/stork/internal/services/albums"
	"github.com/mitsuha/stork/internal/services/artists"
	"github.com/mitsuha/stork/internal/services/overview"
	"github.com/mitsuha/stork/internal/services/playlists"
	"github.com/mitsuha/stork/internal/services/reverseProxy"
	"github.com/mitsuha/stork/internal/services/songs"
	customValidator "github.com/mitsuha/stork/internal/validator"
	"github.com/mitsuha/stork/pkg/authentication"
	"time"
)

func Run() error {
	engine := gin.Default()

	c := cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("audioOnly", customValidator.AudioOnly); err != nil {
			panic(err)
		}
	}

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

	{
		service := artists.New()

		router := r.Group("/artists")
		router.GET("/:id", service.Show)
	}

	{
		service := songs.New()

		router := r.Group("/songs")
		router.GET("/favorite", service.Favorite)
		r.POST("/upload", service.Upload)

		engine.GET("/play/:id", c, service.Play)
	}

	{
		service := playlists.New()

		r.POST("/playlists", service.Create)
		r.GET("/playlists/:id/songs", service.Songs)
		r.PUT("/playlists/:id", service.Update)
	}

	engine.NoRoute(reverseProxy.New())

	return engine.Run()
}
