package handlers

import (
	"crypto/md5"
	"effy/gravatar-profile-card/models"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
	"github.com/gin-gonic/gin"
)

func CreateGravatarURL(email string) (string, string) {
	hash := md5.Sum([]byte(email))
	jsonURL := "https://www.gravatar.com/" + hex.EncodeToString(hash[:]) + ".json"
	imageURL := "https://www.gravatar.com/avatar/" + hex.EncodeToString(hash[:])
	return jsonURL, imageURL
}

func FetchGravatarDetails(jsonURL string) (map[string]interface{}, error) {
	resp, err := http.Get(jsonURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func Gravatar(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Input", "error": err.Error()})
		return
	}

	jsonURL, imageURL := CreateGravatarURL(user.Email)
	jsonDetails, err := FetchGravatarDetails(jsonURL)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch Gravatar details", "error": err.Error()})
		return
	}
	// if err := db.DB.Create(&user).Error; err != nil {
	// 	c.JSON(http.StatusInternalServerError,gin.H{"message":"Failed to store in database"})
	// 	return
	// }

	response := gin.H{
		"image_url":    imageURL,
		"json_url":     jsonURL,
		"json_details": jsonDetails,
		"form_details": user,
	}

	c.JSON(http.StatusOK, response)
}
