package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/LinC3e/shunkan-qr/internal/qr"
)

type QRHandler struct {
	service *qr.Service
}

func NewQRHandler(service *qr.Service) *QRHandler {
	return &QRHandler{service: service}
}

func (h *QRHandler) CreateQR(c *gin.Context) {
	var req struct {
		Content string `json:"content" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	qr, err := h.service.CreateQR(c.Request.Context(), req.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, qr)
}

func (h *QRHandler) GetQR(c *gin.Context) {
	id := c.Param("id")
	qr, err := h.service.GetQR(c.Request.Context(), id)
	if err != nil {
		if err.Error() == "qr not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "QR not found"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, qr)
}