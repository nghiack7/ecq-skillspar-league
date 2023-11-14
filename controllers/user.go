package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nghiack7/ecq-skillspar-league/pkg/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (s *Server) LoginHandler(ctx *gin.Context) {
	// handle Userlogin : 1- username, password, 2 - get database username -> models.User data, 3 compare hashed password and password user -> set ctx data return api_key -> middlware (api_key -> username, passs, role)
	userName, ok := ctx.GetPostForm("user_name")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"error": "user name is not input"})
		return
	}
	Password, ok := ctx.GetPostForm("password")
	if !ok {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "password is not input"})
		return
	}
	var user models.User
	// get data from database of user
	err := s.db.Where("user_name=?", userName).First(&user).Error
	if err != nil || errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.AbortWithStatusJSON(404, gin.H{"error": "user_name is not found"})
		return
	}
	// Comparing the password with the hash
	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(Password))
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": "password is not correct"})
	}
	ctx.JSON(200, gin.H{"api_key": user.ApiKey})
}
