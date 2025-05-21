package http

import (
	"lanchonete/bootstrap"
	"lanchonete/infra/database/mongo"
	"lanchonete/internal/interfaces/http/route"
	"github.com/gin-gonic/gin"
)

// Router handles all HTTP routing for the application
type Router struct {
	env    *bootstrap.Env
	db     mongo.Database
	router *gin.RouterGroup
}

// NewRouter creates a new Router instance
func NewRouter(env *bootstrap.Env, db mongo.Database, ginRouter gin.IRouter) *Router {
	return &Router{
		env:    env,
		db:     db,
		router: ginRouter.Group(""),
	}
}

// Setup initializes all routes for the application
func (r *Router) Setup() {
	// API Documentation
	route.NewDocRouter(r.router)

	// Core domain routes
	route.NewClienteRouter(r.env, r.db, r.router)
	route.NovoPedidoRouter(r.env, r.db, r.router)
	route.NewProdutoRouter(r.env, r.db, r.router)

	// Supporting domain routes
	route.NewPagamentoRouter(r.env, r.db, r.router)
	route.NewAcompanhamentoRouter(r.env, r.db, r.router)
}

// Setup is a package-level function that creates a Router and calls its Setup method
func Setup(env *bootstrap.Env, db mongo.Database, ginRouter gin.IRouter) {
	router := NewRouter(env, db, ginRouter)
	router.Setup()
}