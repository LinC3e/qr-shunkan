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
		api.GET("/qrs", qrHandler.ListQRs)
		api.GET("/qrs/:id", qrHandler.GetQR)
		api.POST("/qrs", qrHandler.CreateQR)
	}

	// PUBLIC QR ROUTE
	r.GET("/qr/:slug", qrHandler.ResolveQR)

	return r
}
