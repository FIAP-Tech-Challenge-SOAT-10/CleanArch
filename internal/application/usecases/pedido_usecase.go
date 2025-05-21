package usecases

import (
	"context"
	"lanchonete/internal/domain/entities"
	"lanchonete/internal/domain/repository"
	"lanchonete/usecases"
)

type pedidoUseCase struct {
	pedidoRepository repository.PedidoRepository
}

func NewPedidoUseCase(pedidoRepository repository.PedidoRepository) usecases.PedidoUseCase {
	return &pedidoUseCase{
		pedidoRepository: pedidoRepository,
	}
}

func (puc *pedidoUseCase) CriarPedido(c context.Context, pedido *entities.Pedido) error {
	return puc.pedidoRepository.CriarPedido(c, pedido)
}

func (puc *pedidoUseCase) BuscarPedido(c context.Context, identificacao string) (*entities.Pedido, error) {
	return puc.pedidoRepository.BuscarPedido(c, identificacao)
}

func (puc *pedidoUseCase) AtualizarStatusPedido(c context.Context, identificacao string, status string, ultimaAtualizacao string) error {
	return puc.pedidoRepository.AtualizarStatusPedido(c, identificacao, status, ultimaAtualizacao)
}
