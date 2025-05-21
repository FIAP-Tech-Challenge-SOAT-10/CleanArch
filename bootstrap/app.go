package bootstrap

import (
	"context"

	"lanchonete/internal/domain/repository"
	"lanchonete/infra/database/mongo"
)

type App struct {
	Env                    *Env
	DB                     mongo.Database
	AcompanhamentoRepository repository.AcompanhamentoRepository
	PedidoRepository         repository.PedidoRepository
	ProdutoRepository        repository.ProdutoRepository
	ClienteRepository        repository.ClienteRepository
	PagamentoRepository      repository.PagamentoRepository
}

func NewApp(ctx context.Context) (*App, error) {
	// Load environment variables
	env := NewEnv()
	
	// Initialize database
	db, err := NewDatabase(ctx, env)
	if err != nil {
		return nil, err
	}

	// Initialize repositories
	acompanhamentoRepo, pedidoRepo, produtoRepo, clienteRepo, pagamentoRepo := NewRepositories(db)

	return &App{
		Env:                    env,
		DB:                     db,
		AcompanhamentoRepository: acompanhamentoRepo,
		PedidoRepository:         pedidoRepo,
		ProdutoRepository:        produtoRepo,
		ClienteRepository:        clienteRepo,
		PagamentoRepository:      pagamentoRepo,
	}, nil
}