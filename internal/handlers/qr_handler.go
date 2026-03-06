package handlers

import (
	"net/http"
	"net/url"
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

	inputURL := c.Query("url")
	sizeParam := c.DefaultQuery("size", "256")
	format := c.DefaultQuery("format", "png")

	if inputURL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "url is required"})
		return
	}

	_, err := url.ParseRequestURI(inputURL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid url"})
		return
	}

	size, err := strconv.Atoi(sizeParam)
	if err != nil {
		size = 256
	}

	data, err := h.service.Generate(inputURL, size, format)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed generating qr"})
		return
	}

	if format == "svg" {
		c.Data(200, "image/svg+xml", data)
		return
	}

	c.Data(200, "image/png", data)
}

func (h *QRHandler) Create(c *gin.Context) {

	var body struct {
		URL string `json:"url"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "invalid body"})
		return
	}

	id := h.service.Create(body.URL)

	c.JSON(200, gin.H{
		"id": id,
		"qr": "/q/" + id,
	})
}

func (h *QRHandler) Resolve(c *gin.Context) {

	id := c.Param("id")

	url, ok := h.service.Get(id)

	if !ok {
		c.JSON(404, gin.H{"error": "not found"})
		return
	}

	c.Redirect(302, url)
}