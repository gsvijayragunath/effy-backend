package handlers

import (
	"crypto/md5"
	"effy/gravatar-profile-card/db"
	"effy/gravatar-profile-card/errors"
	"effy/gravatar-profile-card/models"
	"effy/gravatar-profile-card/utils"
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
	var profile models.Profile
	if err := c.ShouldBindJSON(&profile); err != nil {
		httpStatus, response := utils.RenderError(errors.ErrInvalidRequest, err.Error(), "Invalid Input")
		c.JSON(httpStatus, response)
		// c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Input", "error": err.Error()})
		return
	}

	jsonURL, imageURL := CreateGravatarURL(profile.Email)
	jsonDetails, err := FetchGravatarDetails(jsonURL)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch Gravatar details", "error": err.Error()})
		return
	}
	
	profile.ProfileImageURL = imageURL
	profile.JsonURL = jsonURL

	if err := db.DB.Create(&profile).Error; err != nil {
		httpStatus, response := utils.RenderError(err, err.Error(), "Failed To Create User")
		c.JSON(httpStatus, response)
		return
	}
	response := gin.H{
		"image_url":    imageURL,
		"json_url":     jsonURL,
		"json_details": jsonDetails,
		"form_details": profile,
	}

	c.JSON(http.StatusOK, response)
}
