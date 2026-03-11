package qr

import (
	"net/http"

	"github.com/LinC3e/shunkan-qr/internal/analytics"
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

	// analytics
	ip := c.ClientIP()
	ua := c.Request.UserAgent()

	if h.analytics != nil {
		_ = h.analytics.RegisterScan(
			c.Request.Context(),
			qr.ID.String(),
			ip,
			ua,
		)
	}

	c.JSON(http.StatusOK, qr)
}
