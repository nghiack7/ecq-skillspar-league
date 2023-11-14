package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName       string `json:"user_name"`
	HashedPassword string `json:"hashed_password"`
	ApiKey         string `json: "api_key"`
	Role           Role   `json:"role"`
}

type Role string

const (
	Admin   Role = "admin"
	Player  Role = "player"
	Treaser Role = "treaser"
)
