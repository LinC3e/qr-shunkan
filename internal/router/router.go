package router

import (
	"github.com/LinC3e/shunkan-qr/internal/handlers"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {

	r := gin.Default()

	qrHandler := handlers.NewQRHandler()

	api := r.Group("/api")

	api.GET("/qr", qrHandler.Generate)
	api.POST("/qr", qrHandler.Create)

	r.GET("/q/:id", qrHandler.Resolve)

	r.Static("/static", "./web/static")

	r.GET("/", func(c *gin.Context) {
		c.File("./web/index.html")
	})

	return r
}