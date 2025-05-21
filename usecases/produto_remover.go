package usecases

import (
	"context"
	"fmt"
	"lanchonete/internal/domain/repository"
)

type ProdutoRemoverUseCase interface {
	Run(ctx context.Context, identificacao string) error
}

type produtoRemoverUseCase struct {
	produtoGateway repository.ProdutoRepository
}

func NewProdutoRemoverUseCase(produtoGateway repository.ProdutoRepository) ProdutoRemoverUseCase {
	return &produtoRemoverUseCase{
		produtoGateway: produtoGateway,
	}
}

func (pruc *produtoRemoverUseCase) Run(c context.Context, identificacao string) error {
	_, err := pruc.produtoGateway.BuscarProdutoPorId(c, identificacao)
	if err != nil {
		return fmt.Errorf("produto não existe no banco de dados: %w", err)
	}

	err = pruc.produtoGateway.RemoverProduto(c, identificacao)
	if err != nil {
		return fmt.Errorf("não foi possível remover o produto: %w", err)
	}

	return nil
}
