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
	Data string `json:"data"` // Input field for QR code data
}

func (h *qrcodehandler) GenerateQR(c *gin.Context) {
	var input QrInput

	// Bind the input JSON to the QrInput struct
	err := c.ShouldBindJSON(&input)
	if err != nil {
		httpStatus, response := utils.RenderError(errors.ErrInvalidRequest, err.Error(), "Invalid Input")
		c.JSON(httpStatus, response)
		return
	}

	// Generate the QR Code
	qrCode, err := qrcode.Encode(input.Data, qrcode.Medium, 256)
	if err != nil {
		httpStatus, response := utils.RenderError(err, err.Error(), "Failed to generate QR code")
		c.JSON(httpStatus, response)
		return
	}

	// Convert the QR Code to Base64
	base64QRCode := base64.StdEncoding.EncodeToString(qrCode)

	// Send the response with the Base64 QR Code
	c.JSON(http.StatusOK, gin.H{
		"message": "QR Code generated successfully",
		"qr_code": base64QRCode,
	})
}
