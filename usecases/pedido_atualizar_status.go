package usecases

import (
	"context"
	"lanchonete/internal/domain/entities"
	"lanchonete/internal/domain/repository"
)

type PedidoAtualizarStatusUseCase interface {
	Run(ctx context.Context, identificacao string, status string) error
}

type pedidoAtualizarStatusUseCase struct {
	pedidoGateway repository.PedidoRepository
}

func NewPedidoAtualizarStatusUseCase(pedidoGateway repository.PedidoRepository) PedidoAtualizarStatusUseCase {
	return &pedidoAtualizarStatusUseCase{
		pedidoGateway: pedidoGateway,
	}
}

func (pduc *pedidoAtualizarStatusUseCase) Run(c context.Context, identificacao string, status string) error {

	pedido, err := pduc.pedidoGateway.BuscarPedido(c, identificacao)
	if err != nil {
		return err
	}

	timeStamp, err := pedido.UpdateStatus(entities.StatusPedido(status))
	if err != nil {
		return err
	}
	err = pduc.pedidoGateway.AtualizarStatusPedido(c, identificacao, status, timeStamp)
	if err != nil {
		return err
	}
	return nil
}
