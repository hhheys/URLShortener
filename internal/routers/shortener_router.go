package routers

import (
	"URLShortener/internal/app"

	"github.com/gin-gonic/gin"
)

func NewShortenerRouter(r *gin.Engine, app *app.App) {
	g := r.Group("")

	g.POST("/make_short/", app.Handlers.ShortenerHandler.MakeShortURL)
	g.GET("/:url", app.Handlers.ShortenerHandler.UseShortURL)
}
