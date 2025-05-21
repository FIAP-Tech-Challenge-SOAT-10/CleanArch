package route

import (
	handler "lanchonete/internal/interfaces/http/handlers"
	"lanchonete/internal/application/usecases"
	"lanchonete/bootstrap"
	"lanchonete/infra/database/mongo"
	"lanchonete/internal/infrastructure/repository"

	"github.com/gin-gonic/gin"
)

// NewPagamentoRouter creates and configures all pagamento-related routes
func NewPagamentoRouter(env *bootstrap.Env, db mongo.Database, router *gin.RouterGroup) {

	// Create repository using the repository factory
	repositorioFactory := repository.NovoRepositorioFactory(db)
	pagamentoRepo := repositorioFactory.CriarPagamentoRepository()
	
	pc := &handler.PagamentoHandler{
		EnviarPagamentoUseCase: usecases.NewEnviarPagamentoUseCase(pagamentoRepo),
		ConfirmarPagamentoUseCase: usecases.NewConfirmarPagamentoUseCase(pagamentoRepo),
	}

	// Register payment routes
	router.POST("/pagamento", pc.EnviarPagamento)
	router.POST("/pagamento/confirmar", pc.ConfirmarPagamento)
}