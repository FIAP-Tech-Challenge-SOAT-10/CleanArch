package usecases

import (
	"context"
	"fmt"
	"lanchonete/internal/domain/entities"
	"lanchonete/internal/domain/repository"
)

type ProdutoEditarUseCase interface {
	Run(ctx context.Context, identificacao, nome, categoria, descricao string, preco float32) (*entities.Produto, error)
}

type produtoEditarUseCase struct {
	produtoGateway repository.ProdutoRepository
}

func NewProdutoEditarUseCase(produtoGateway repository.ProdutoRepository) ProdutoEditarUseCase {
	return &produtoEditarUseCase{
		produtoGateway: produtoGateway,
	}
}

func (puc *produtoEditarUseCase) Run(c context.Context, identificacao string, nome string, categoria string, descricao string, preco float32) (*entities.Produto, error) {

	produto, err := puc.produtoGateway.BuscarProdutoPorId(c, identificacao)

	if err != nil {
		return nil, fmt.Errorf("produto não cadastrado, crie o produto primeiro: %w", err)
	}

	if nome == "" {
		nome = produto.Nome
	}

	if categoria == "" {
		categoria = string(produto.Categoria)
	}

	if descricao == "" {
		descricao = produto.Descricao
	}

	if preco == 0 {
		preco = produto.Preco
	}

	produtoEditado, err := entities.ProdutoNew(identificacao, nome, categoria, descricao, preco)
	if err != nil {
		return nil, fmt.Errorf("atualização de produto inválida: %w", err)
	}

	err = puc.produtoGateway.EditarProduto(c, produtoEditado)
	if err != nil {
		return nil, fmt.Errorf("não foi possível atualizar o produto: %w", err)

	}

	return produto, nil
}
