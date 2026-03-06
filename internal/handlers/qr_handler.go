package handlers

import (
	"net/http"
	"strconv"

	"github.com/LinC3e/shunkan-qr/internal/qr"

	"github.com/gin-gonic/gin"
)

type QRHandler struct {
	service *qr.Service
}

func NewQRHandler() *QRHandler {
	return &QRHandler{
		service: qr.NewService(),
	}
}

func (h *QRHandler) Generate(c *gin.Context) {

	url := c.Query("url")
	sizeParam := c.Query("size")

	if url == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "url is required",
		})
		return
	}

	size := 256

	if sizeParam != "" {
		s, err := strconv.Atoi(sizeParam)
		if err == nil {
			size = s
		}
	}

	png, err := h.service.Generate(url, size)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed generating qr",
		})
		return
	}

	c.Data(http.StatusOK, "image/png", png)

}