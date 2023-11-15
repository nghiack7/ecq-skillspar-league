package models

type Fund struct {
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
