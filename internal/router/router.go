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
		// QR
		api.GET("/qrs", qrHandler.ListQRs)
		api.GET("/qrs/:id", qrHandler.GetQR)
		api.POST("/qrs", qrHandler.CreateQR)

		// QR analytics
		api.GET("/qrs/:id/stats", qrHandler.GetStats)
		api.GET("/qrs/:id/stats/daily", qrHandler.GetDailyStats)
		/* 		api.GET("/qrs/:id/stats/countries", qrHandler.GetCountryStats)
		   		api.GET("/qrs/:id/stats/devices", qrHandler.GetDeviceStats)
		   		api.GET("/qrs/:id/stats/browsers", qrHandler.GetBrowserStats) */
	}

	// PUBLIC QR ROUTE
	r.GET("/qr/:slug", qrHandler.ResolveQR)

	return r
}
