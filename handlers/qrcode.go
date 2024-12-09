package handlers

import (
	"effy/gravatar-profile-card/errors"
	"effy/gravatar-profile-card/utils"
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

type qrcodehandler struct{}

func NewQrCodehandler() *qrcodehandler {
	return &qrcodehandler{}
}

type QrInput struct {
	Data string `json:"data"` 
}

func (h *qrcodehandler) GenerateQR(c *gin.Context) {
	var input QrInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		httpStatus, response := utils.RenderError(errors.ErrInvalidRequest, err.Error(), "Invalid Input")
		c.JSON(httpStatus, response)
		return
	}

	qrCode, err := qrcode.Encode(input.Data, qrcode.Medium, 256)
	if err != nil {
		httpStatus, response := utils.RenderError(err, err.Error(), "Failed to generate QR code")
		c.JSON(httpStatus, response)
		return
	}

	base64QRCode := base64.StdEncoding.EncodeToString(qrCode)

	c.JSON(http.StatusOK, gin.H{
		"message": "QR Code generated successfully",
		"qr_code": base64QRCode,
	})
}
