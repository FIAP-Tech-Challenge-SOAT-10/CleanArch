package server

import (
	"lanchonete/bootstrap"
	"lanchonete/internal/interfaces/http"
	"lanchonete/infra/database/mongo"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	env    *bootstrap.Env
	db     mongo.Database
	router *gin.Engine
}

func NewServer(env *bootstrap.Env, db mongo.Database) *Server {
	router := gin.Default()
	
	// Configure CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	return &Server{
		env:    env,
		db:     db,
		router: router,
	}
}

func (s *Server) SetupRoutes() {
	http.Setup(s.env, s.db, s.router)
}

func (s *Server) Start() error {
	return s.router.Run(s.env.ServerAddress)
} 