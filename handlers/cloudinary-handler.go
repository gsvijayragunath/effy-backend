package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
)

type chandler struct{}

func NewChandler() *chandler {
	return &chandler{}
}

func (h *chandler) UploadAndGeneratePublicURL(c *gin.Context) {
	cld, err := cloudinary.NewFromParams("ddx5mjhm6", "285947226752892", "YyvVuviX5MPCbJ7QTe4NwUQ9e7k")
	if err != nil {
		log.Printf("Cloudinary initialization error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to initialize Cloudinary"})
		return
	}

	// Parse the uploaded file
	file, err := c.FormFile("image") // "image" is the key for the uploaded file
	if err != nil {
		log.Printf("File parsing error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse file"})
		return
	}

	// Open the file for uploading
	uploadedFile, err := file.Open()
	if err != nil {
		log.Printf("File opening error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to open file"})
		return
	}
	defer uploadedFile.Close()

	// Upload the file to Cloudinary
	uploadResult, err := cld.Upload.Upload(context.Background(), uploadedFile, uploader.UploadParams{
		Folder: "profile_cards", // Optional folder name in Cloudinary
	})
	if err != nil {
		log.Printf("File upload error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file"})
		return
	}

	// Return the public URL
	publicURL := uploadResult.SecureURL
	c.JSON(http.StatusOK, gin.H{
		"message":   "File uploaded successfully",
		"publicURL": publicURL,
	})
}
