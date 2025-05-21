package route

import (
	"lanchonete/bootstrap"
	"lanchonete/infra/database/mongo"
	"lanchonete/internal/infrastructure/repository"
	handler "lanchonete/internal/interfaces/http/handlers"
	"lanchonete/usecases"

	"github.com/gin-gonic/gin"
)

// NewProdutoRouter creates and configures all produto-related routes
func NewProdutoRouter(env *bootstrap.Env, db mongo.Database, router *gin.RouterGroup) {
	// Create repository using the repository factory
	repositorioFactory := repository.NovoRepositorioFactory(db)
	produtoRepo := repositorioFactory.CriarProdutoRepository()

	pc := &handler.ProdutoHandler{
		ProdutoIncluirUseCase:            usecases.NewProdutoIncluirUseCase(produtoRepo),
		ProdutoBuscarPorIdUseCase:        usecases.NewProdutoBuscaPorIdUseCase(produtoRepo),
		ProdutoListarTodosUseCase:        usecases.NewProdutoListarTodosUseCase(produtoRepo),
		ProdutoEditarUseCase:             usecases.NewProdutoEditarUseCase(produtoRepo),
		ProdutoRemoverUseCase:            usecases.NewProdutoRemoverUseCase(produtoRepo),
		ProdutoListarPorCategoriaUseCase: usecases.NewProdutoListarPorCategoriaUseCase(produtoRepo),
	}

	router.POST("/produto", pc.ProdutoIncluir)
	router.GET("/produto/:id", pc.ProdutoBuscarPorId)
	router.GET("/produtos", pc.ProdutoListarTodos)
	router.GET("/produtos/:categoria", pc.ProdutoListarPorCategoria)
	router.POST("/produto/editar", pc.ProdutoEditar)
	router.DELETE("/produto/:id", pc.ProdutoRemover)
}
