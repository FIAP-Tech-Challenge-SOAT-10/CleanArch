package repository

import (
	"context"
	"lanchonete/internal/domain/entities"
)

// PedidoRepository define a interface para operações de dados de pedidos
type PedidoRepository interface {
	CriarPedido(c context.Context, pedido *entities.Pedido) error
	BuscarPedido(c context.Context, identificacao string) (*entities.Pedido, error)
	AtualizarStatusPedido(c context.Context, Identificacao string, status string, UltimaAtualizacao string) error
	ListarTodosOsPedidos(c context.Context) ([]*entities.Pedido, error)
}
