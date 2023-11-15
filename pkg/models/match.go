package models

import (
	"fmt"

	"gorm.io/gorm"
)

type FootBallPlayer struct {
	gorm.Model
	PlayerName string `json:"player_name"`
	UserID     uint   `-`
	Join       bool   `json:"join"`
	Team       int    `json:"team"`
}

func RegisterMatch(user User) error {
	var player FootBallPlayer
	player.UserID = user.ID
	player.PlayerName = user.FullName
	player.Join = true
	err := dbGorm.Save(&player).Error
	if err != nil {
		return err
	}
	return nil
}

func CancelAllUser() error {
	var players []FootBallPlayer
	err := dbGorm.Find(&players).Error
	if err != nil {
		return err
	}
	for i := range players {
		players[i].Join = false
		players[i].Team = 0
	}
	err = dbGorm.Save(&players).Error
	if err != nil {
		return err
	}
	return nil
}

func Cancel(u User) error {
	var player FootBallPlayer
	err := dbGorm.Where("user_id=?", u.ID).First(&player).Error
	if err != nil {
		return err
	}
	if !player.Join {
		return fmt.Errorf("user is not register for this match")
	}
	player.Join = false
	player.Team = 0
	err = dbGorm.Save(&player).Error
	if err != nil {
		return err
	}
	return nil
}

func ListRegisterFootball() []FootBallPlayer {
	var players []FootBallPlayer
	err := dbGorm.Where("join=?", true).Find(&players).Error
	if err != nil {
		return nil
	}
	return players
}
func listTeamPlayers() (map[int][]FootBallPlayer, error) {
	teamPlayer := make(map[int][]FootBallPlayer)
	var players []FootBallPlayer
	err := dbGorm.Where("join=?", true).Find(&players).Error
	if err != nil {
		return nil, err
	}
	for _, player := range players {
		teamPlayer[player.Team] = append(teamPlayer[player.Team], player)
	}
	return teamPlayer, nil
}

func ResultMatch(rank map[int]int, totalCredits int64) error {
	teamPlayer, err := listTeamPlayers()
	if err != nil {
		return err
	}
	for k, v := range rank {
		err = updateResultMatch(teamPlayer[k], v)
		if err != nil {
			return err
		}
		err = updateFundMatch(teamPlayer[k], v, totalCredits)
		if err != nil {
			return err
		}
	}
	go CancelAllUser()
	return nil
}
