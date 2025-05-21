package route

import (
	"lanchonete/bootstrap"
	"lanchonete/infra/database/mongo"
	factory "lanchonete/internal/application/factories"
	appusecases "lanchonete/internal/application/usecases"
	"lanchonete/internal/infrastructure/repository"
	handler "lanchonete/internal/interfaces/http/handlers"

	"fmt"

	"github.com/gin-gonic/gin"
)

// NewAcompanhamentoRouter creates and configures all acompanhamento-related routes
func NewAcompanhamentoRouter(env *bootstrap.Env, db mongo.Database, router *gin.RouterGroup) {
	// Criar fábricas
	repositorioFactory := repository.NovoRepositorioFactory(db)
	useCaseFactory := factory.NovoUseCaseFactory(repositorioFactory)

	// Criar casos de uso
	acompanhamentoUseCase := appusecases.NewAcompanhamentoUseCase(repositorioFactory.CriarAcompanhamentoRepository())
	pedidoAtualizarStatusUseCase := useCaseFactory.CriarPedidoAtualizarStatusUseCase()

	auc := &handler.AcompanhamentoHandler{
		AcompanhamentoUseCase:        acompanhamentoUseCase,
		PedidoAtualizarStatusUseCase: pedidoAtualizarStatusUseCase,
	}

	fmt.Printf("Registrando rotas do acompanhamento\n")

	router.POST("/acompanhamento", auc.CriarAcompanhamento)
	router.GET("/acompanhamento/show", auc.BuscarAcompanhamento)
	router.GET("/acompanhamento/:ID", auc.BuscarPedido)
	router.POST("/acompanhamento/:IDAcompanhamento/:IDPedido", auc.AdicionarPedido)
	router.PUT("acompanhamento/:IDAcompanhamento/:IDPedido/:status", auc.AtualizarStatusPedido)
}
