package route

import (
	handler "lanchonete/internal/interfaces/http/handlers"
	"lanchonete/bootstrap"
	"lanchonete/internal/infrastructure/repository"
	"lanchonete/infra/database/mongo"
	"lanchonete/internal/application/usecases"

	"github.com/gin-gonic/gin"
)

// NewClienteRouter creates and configures all cliente-related routes
func NewClienteRouter(env *bootstrap.Env, db mongo.Database, router *gin.RouterGroup) {
	// Criar fábricas
	repositorioFactory := repository.NovoRepositorioFactory(db)
	
	// Criar casos de uso
	clienteUseCase := usecases.NewClienteUseCase(repositorioFactory.CriarClienteRepository())
	
	cc := &handler.ClienteHandler{
		ClienteUseCase: clienteUseCase,
	}

	router.GET("/cliente/:CPF", cc.BuscarCliente)
	router.POST("/cliente", cc.CriarCliente)
}