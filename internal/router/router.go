package router

import (
	"github.com/gin-gonic/gin"

	handler "github.com/LinC3e/shunkan-qr/internal/handlers"
)

func Setup(qrHandler *handler.QRHandler) *gin.Engine {
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