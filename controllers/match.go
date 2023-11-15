package controllers

import (
	"fmt"
	"strconv"
	"strings"

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

func (s *Server) UpdateResultMatch(c *gin.Context) {
	rankTeamm, ok := c.GetPostForm("rank")
	if !ok {
		c.AbortWithStatusJSON(400, gin.H{"message": "input rank team should be a number"})
		return
	}
	credit, ok := c.GetPostForm("total_credits")
	if !ok {
		c.AbortWithStatus(400)
		return
	}
	totalCredit, _ := strconv.ParseInt(credit, 10, 64)
	team := strings.Split(rankTeamm, ",")
	rank := make(map[int]int)
	for i, t := range team {
		j, _ := strconv.ParseInt(t, 10, 64)
		rank[int(j)] = i + 1
	}
	err := models.ResultMatch(rank, totalCredit)
	if err != nil {
		c.AbortWithError(500, err)
	}
	c.JSON(200, gin.H{"message": "update data match successfully"})
}
