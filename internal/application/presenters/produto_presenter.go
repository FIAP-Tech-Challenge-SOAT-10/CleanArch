package presenters

import "lanchonete/internal/domain/entities"

type ProdutoDTO struct {
    Identificacao 	string `json:"identificacao"`
    Nome			string `json:"nome"`
	Categoria   	entities.CatProduto `json:"categoria"`
	Descricao		string `json:"descricao"`
	Preco			float32 `json:"preco"` 
}

func NewProdutoDTO(produto *entities.Produto) (*ProdutoDTO) {
    if produto == nil {
        return nil
    }
    
    return &ProdutoDTO{
        Identificacao: produto.Identificacao,
        Nome: 		produto.Nome,
		Categoria:  produto.Categoria,
		Descricao:  produto.Descricao,
		Preco: 		produto.Preco,
    }
}