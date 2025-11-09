package routers

import (
	"URLShortener/internal/app"

	"github.com/gin-gonic/gin"
)

func NewRouter(app *app.App) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())

	NewShortenerRouter(r, app)

	return r
}
