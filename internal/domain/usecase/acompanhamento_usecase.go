package usecase

import (
	"context"
	"lanchonete/internal/domain/entities"
)

// AcompanhamentoUseCase define a interface para operações de acompanhamento de pedidos
type AcompanhamentoUseCase interface {
	CriarAcompanhamento(c context.Context, acompanhamento *entities.AcompanhamentoPedido) error
	BuscarPedido(c context.Context, ID string) (entities.Pedido, error)
	AdicionarPedido(c context.Context, acompanhamentoID string, pedido *entities.Pedido) error
	BuscarAcompanhamento(c context.Context, ID string) (*entities.AcompanhamentoPedido, error)
	AtualizarStatusPedido(c context.Context, acompanhamentoID string, identificacao string, novoStatus entities.StatusPedido) error
} 