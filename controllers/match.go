package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nghiack7/ecq-skillspar-league/pkg/models"
)

func (s *Server) RegisterMatch(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		c.Abort()
		return
	}
	userStr := user.(models.User)
	err := models.RegisterMatch(userStr)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(200, gin.H{userStr.FullName: "Register Success Fully For This Match"})
}

func (s *Server) CancelAllUser(c *gin.Context) {
	err := models.CancelAllUser()
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(200, gin.H{"message": "reset all data from user"})
}

func (s *Server) Cancel(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		c.Abort()
		return
	}
	userStr := user.(models.User)
	err := models.Cancel(userStr)
	if err != nil {
		c.AbortWithError(404, err)
	}
	c.JSON(200, gin.H{"message": fmt.Sprintf("user %s cancel successfully for this match", userStr.FullName)})
}

func (s *Server) GetListFootball(c *gin.Context) {
	players := models.ListRegisterFootball()
	c.JSON(200, players)
}
