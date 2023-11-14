package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName       string `json:"user_name"`
	FullName       string `json:"full_name"`
	HashedPassword string `json:"hashed_password"`
	ApiKey         string `json:"api_key"`
	Role           Role   `json:"role"`
}

type Role string

const (
	Admin   Role = "admin"
	Player  Role = "player"
	Treaser Role = "treaser"
)

func GetUserByAPIKey(apiKey string) (User, error) {
	var user User
	err := dbGorm.Where("api_key=?", apiKey).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
