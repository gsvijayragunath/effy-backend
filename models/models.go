package models

import "github.com/google/uuid"

type Profile struct {
	ProfileCardID   uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"profilecard_id"`
	Email           string `json:"email" binding:"required"`
	FullName        string `json:"full_name" binding:"required"`
	Username        string `json:"user_name" binding:"required"`
	PhoneNumber     string `json:"phone_number" binding:"required"`
	Location        string `json:"location" binding:"required"`
	Bio             string `json:"bio" binding:"required"`
	Website         string `json:"website"`
	ProfileImageURL string `json:"profile_image"`
	JsonURL         string `json:"json_url"`
}

type User struct {
	UserID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"user_id"`
	Name            string    `json:"name" binding:"required"`
	Email           string    `gorm:"uniqueIndex;not null" json:"email" binding:"required"`
	Country         string    `json:"country" binding:"required"`
	UserType        string    `json:"user_type" binding:"required"`
	Password        string    `json:"password" binding:"required"`
}

type Signin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
