package models

import "gorm.io/gorm"

type Fund struct {
	gorm.Model
	UserID uint
	Credit int64
}

func AddCredit(userID uint, money int64) error {
	var fund Fund
	tx := dbGorm.Begin()
	err := tx.Where("user_id=?", userID).First(&fund).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	fund.Credit += money
	err = tx.Save(&fund).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func UsedCredit(userID uint, money int64) error {
	var fund Fund
	tx := dbGorm.Begin()
	err := tx.Where("user_id=?", userID).First(&fund).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	fund.Credit -= money
	err = tx.Save(&fund).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func updateFundMatch(players []FootBallPlayer, rank int, totalCredits int64) error {
	num := len(players)
	money2 := totalCredits * 3 / 10
	money3 := totalCredits - money2
	var funds []Fund
	err := dbGorm.Find(&funds).Error
	if err != nil {
		return err
	}
	mapFund := make(map[uint]Fund)
	for _, f := range funds {
		mapFund[f.UserID] = f
	}
	var eachCredit int64
	switch rank {
	case 1:
		return nil
	case 2:
		eachCredit = money2 / int64(num)
	case 3:
		eachCredit = money3 / int64(num)
	}
	var newfunds []Fund

	for _, p := range players {
		fund := mapFund[p.UserID]
		fund.Credit -= eachCredit
		newfunds = append(newfunds, fund)
	}
	err = dbGorm.Save(&newfunds).Error
	if err != nil {
		return err
	}
	return nil
}
