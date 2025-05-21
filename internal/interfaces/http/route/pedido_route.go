package route

import (
	"lanchonete/bootstrap"
	"lanchonete/infra/database/mongo"
	factory "lanchonete/internal/application/factories"
	"lanchonete/internal/infrastructure/repository"
	handler "lanchonete/internal/interfaces/http/handlers"

	"github.com/gin-gonic/gin"
)

// NovoPedidoRouter cria um novo roteador para pedidos
func NovoPedidoRouter(env *bootstrap.Env, db mongo.Database, router *gin.RouterGroup) {
	// Criar fábricas
	repositorioFactory := repository.NovoRepositorioFactory(db)
	useCaseFactory := factory.NovoUseCaseFactory(repositorioFactory)

	// Criar casos de uso
	pedidoIncluirUseCase := useCaseFactory.CriarPedidoIncluirUseCase()
	pedidoBuscarPorIdUseCase := useCaseFactory.CriarPedidoBuscarPorIdUseCase()
	pedidoAtualizarStatusUseCase := useCaseFactory.CriarPedidoAtualizarStatusUseCase()
	produtoBuscarPorIdUseCase := useCaseFactory.CriarProdutoBuscarPorIdUseCase()
	pedidoListarTodosUseCase := useCaseFactory.CriarPedidoListarTodosUseCase()

	// Criar handler
	puc := &handler.PedidoHandler{
		PedidoIncluirUseCase:         pedidoIncluirUseCase,
		PedidoBuscarPorIdUseCase:     pedidoBuscarPorIdUseCase,
		PedidoAtualizarStatusUseCase: pedidoAtualizarStatusUseCase,
		ProdutoBuscarPorIdUseCase:    produtoBuscarPorIdUseCase,
		PedidoListarTodosUseCase:     pedidoListarTodosUseCase,
	}

	// Configurar rotas
	router.POST("/pedidos", puc.CriarPedido)
	router.GET("/pedidos/:nroPedido", puc.BuscarPedido)
	router.PUT("/pedidos/:nroPedido/status/:status", puc.AtualizarStatusPedido)
	router.POST("/pedidos/listartodos", puc.ListarTodosOsPedidos)
}
