package models


type User struct {
	Email       string `json:"email" binding:"required"`
	FullName    string `json:"full_name" binding:"required"`
	Username    string `json:"user_name" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Location    string `json:"location" binding:"required"`
	Bio         string `json:"bio" binding:"required"`
	Website     string `json:"website"`
	ProfileImageURL string `json:"profile_image"`
	JsonURL string `json:"json_url"`
}
