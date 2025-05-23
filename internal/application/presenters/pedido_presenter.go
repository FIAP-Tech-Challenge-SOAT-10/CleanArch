package presenters

import (
	"lanchonete/internal/domain/entities"
	"time"
)

// PedidoDTO representa os dados de um pedido para apresentação
type PedidoDTO struct {
	ID            string                `json:"id"`
	Identificacao string                `json:"identificacao"`
	Status        entities.StatusPedido `json:"status"`
	TempoEstimado time.Duration         `json:"tempoEstimado"`
	Itens         []ItemPedidoDTO       `json:"itens"`
	Cliente       ClienteDTO            `json:"cliente"`
	Total         float32               `json:"total"`
}

// ItemPedidoDTO representa os dados de um item de pedido para apresentação
type ItemPedidoDTO struct {
	ProdutoID     string  `json:"produtoId"`
	NomeProduto   string  `json:"nomeProduto"`
	Quantidade    int     `json:"quantidade"`
	PrecoUnitario float32 `json:"precoUnitario"`
	Subtotal      float32 `json:"subtotal"`
}

// NewPedidoDTO cria um novo DTO a partir de uma entidade Pedido
func NewPedidoDTO(p *entities.Pedido) *PedidoDTO {
	itens := make([]ItemPedidoDTO, 0)
	for _, produto := range p.Produtos {
		itens = append(itens, ItemPedidoDTO{
			ProdutoID:     produto.Identificacao,
			NomeProduto:   produto.Nome,
			Quantidade:    1, // Default quantity since there's no quantity field in Produto
			PrecoUnitario: produto.Preco,
			Subtotal:      produto.Preco,
		})
	}

	return &PedidoDTO{
		ID:            p.Identificacao, 
		Identificacao: p.Identificacao,
		Status:        p.Status,
		TempoEstimado: time.Duration(900), // Default 15 minutes
		Itens:         itens,
		Cliente: ClienteDTO{
			Nome:  p.Cliente.Nome,
			CPF:   p.Cliente.CPF,
			Email: p.Cliente.Email,
		},
		Total: p.Total,
	}
} 