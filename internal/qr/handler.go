package qr

import (
	"net/http"

	"github.com/LinC3e/shunkan-qr/internal/analytics"
	"github.com/LinC3e/shunkan-qr/internal/utils"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service   *Service
	analytics *analytics.Service
}

func NewHandler(service *Service, analytics *analytics.Service) *Handler {
	return &Handler{
		service:   service,
		analytics: analytics,
	}
}

type CreateQRRequest struct {
	Content string `json:"content" binding:"required"`
}

func (h *Handler) ListQRs(c *gin.Context) {

	qrs, err := h.service.ListQRs(c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, qrs)
}

func (h *Handler) CreateQR(c *gin.Context) {

	var req CreateQRRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	qr, err := h.service.CreateQR(c.Request.Context(), req.Content)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, qr)
}

func (h *Handler) GetQR(c *gin.Context) {

	id := c.Param("id")

	qr, err := h.service.GetQR(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "qr not found",
		})
		return
	}

	c.JSON(http.StatusOK, qr)
}

func (h *Handler) ResolveQR(c *gin.Context) {

	slug := c.Param("slug")

	qr, err := h.service.ResolveQR(c.Request.Context(), slug)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "qr not found",
		})
		return
	}

	ip := c.ClientIP()
	ua := c.Request.UserAgent()

	// utils
	device, browser := utils.ParseUserAgent(ua)
	country := utils.GetCountry(ip)

	if h.analytics != nil {
		_ = h.analytics.RegisterScan(
			c.Request.Context(),
			qr.ID.String(),
			ip,
			ua,
			country,
			device,
			browser,
		)
	}

	c.JSON(http.StatusOK, qr)
}

func (h *Handler) GetStats(c *gin.Context) {

	id := c.Param("id")

	stats, err := h.analytics.GetStats(c.Request.Context(), id)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to get stats"})
		return
	}

	c.JSON(200, stats)
}

func (h *Handler) GetDailyStats(c *gin.Context) {

	id := c.Param("id")

	stats, err := h.analytics.GetDailyStats(c.Request.Context(), id)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to get daily stats"})
		return
	}

	c.JSON(200, stats)
}
