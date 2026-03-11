package router

import (
	"github.com/gin-gonic/gin"

	"github.com/LinC3e/shunkan-qr/internal/qr"
)

func Setup(qrHandler *qr.Handler) *gin.Engine {
	r := gin.Default()

	r.Static("/static", "./web/static")
	r.StaticFile("/", "./web/index.html")

	api := r.Group("/api")
	{
		api.POST("/qrs", qrHandler.CreateQR)
		api.GET("/qrs/:id", qrHandler.GetQR)
	}

	return r
}