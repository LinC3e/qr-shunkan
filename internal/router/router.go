package router

import (
	"github.com/LinC3e/shunkan-qr/internal/handlers"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {

	r := gin.Default()

	qrHandler := handlers.NewQRHandler()

	r.GET("/api/qr", qrHandler.Generate)
	r.Static("/static", "./web/static")

	// index
	r.GET("/", func(c *gin.Context) {
		c.File("./web/index.html")
	})

	return r
}