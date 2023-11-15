package models

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName       string   `json:"user_name"`
	FullName       string   `json:"full_name"`
	HashedPassword string   `json:"hashed_password"`
	ApiKey         string   `json:"api_key"`
	Role           Role     `json:"role"`
	Played         int64    `json:"played"`
	Position       Position `json:"position"`
	Point          int64    `json:"point"`
	Rate           string   `json:"rate"`
}

type Position struct {
	Top    int64 `json:"top"`
	Middle int64 `json:"middle"`
	Bottom int64 `json:"bottom"`
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

func updateResultMatch(players []FootBallPlayer, rank int) error {
	for _, player := range players {
		var user User
		err := dbGorm.Where("id=?", player.ID).First(&user).Error
		if err != nil {
			return err
		}
		switch rank {
		case 1:
			user.Position.Top += 1
		case 2:
			user.Position.Middle += 1
		case 3:
			user.Position.Bottom += 1
		}
		user.Played = user.Position.Top + user.Position.Middle + user.Position.Bottom
		user.Point = user.Position.Top*3 + user.Position.Middle*2 + user.Position.Bottom
		user.Rate = fmt.Sprintf("%.2f", float64(user.Point)/float64(user.Played)) // Formatting Rate field

		err = dbGorm.Save(&user).Error
		if err != nil {
			return err
		}
	}
	return nil
}
