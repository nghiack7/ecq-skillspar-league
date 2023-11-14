package models

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite" // Sqlite driver based on CGO

	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"gorm.io/gorm"
)

func InitDatabase() (*gorm.DB, error) {

	fmt.Println("init database success")
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(User{})
	adimUser := User{
		UserName: "nghiack7",
		Role:     "admin",
	}
	password := "admin"
	hashedPass, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	adimUser.HashedPassword = string(hashedPass)
	adimUser.ApiKey = "abaklsjvklasvklasm"
	err = db.Save(&adimUser).Error
	if err != nil {
		log.Fatal(err)
	}
	return db, err
	// TODO: Init database and connect to database
}
