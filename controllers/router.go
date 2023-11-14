package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nghiack7/ecq-skillspar-league/middlewares"
	"gorm.io/gorm"
)

type Server struct {
	svr *http.Server
	db  *gorm.DB
}

func (s *Server) StartServer() {
	log.Fatal(s.svr.ListenAndServe())
}

func InitServer(db *gorm.DB) *Server {
	r := gin.Default()
	s := &Server{svr: &http.Server{}, db: db}
	s.svr.Handler = r
	s.svr.Addr = "localhost:8080"
	s.initRoute(r)

	return s
}

func (s *Server) initRoute(r *gin.Engine) {
	r.POST("/login", s.LoginHandler)
	api := r.Group("/api")
	api.Use(middlewares.RequiredApiKey())
	api.GET("/match", s.GetListFootball)
	api.POST("/register", s.RegisterMatch)
	api.POST("/cancel", s.Cancel)
}
