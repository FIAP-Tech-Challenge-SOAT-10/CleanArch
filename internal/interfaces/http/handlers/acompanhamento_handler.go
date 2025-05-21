package handler

import (
	"fmt"
	_ "lanchonete/docs"
	"lanchonete/internal/domain/entities"
	"lanchonete/internal/domain/usecase"
	"lanchonete/usecases"
	response "lanchonete/internal/interfaces/http/responses"
	"lanchonete/internal/application/presenters"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AcompanhamentoHandler struct {
	AcompanhamentoUseCase        usecase.AcompanhamentoUseCase
	PedidoAtualizarStatusUseCase usecases.PedidoAtualizarStatusUseCase
}

func NewAcompanhamentoHandler(auc usecase.AcompanhamentoUseCase, p usecases.PedidoAtualizarStatusUseCase) *AcompanhamentoHandler {
	return &AcompanhamentoHandler{
		AcompanhamentoUseCase:        auc,
		PedidoAtualizarStatusUseCase: p,
	}
}

// CriarAcompanhamento godoc
// @Summary Cria um acompanhamento
// @Description Cria um acompanhamento
// @Tags acompanhamento
// @Router /acompanhamento [post]
// @Accept  json
// @Produce  json
// @Param acompanhamento body entities.AcompanhamentoPedido true "Acompanhamento"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
func (ah *AcompanhamentoHandler) CriarAcompanhamento(a *gin.Context) {
	fmt.Print(a)
	var acompanhamento entities.AcompanhamentoPedido

	err := a.ShouldBind(&acompanhamento)
	if err != nil {
		a.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	err = ah.AcompanhamentoUseCase.CriarAcompanhamento(a, &acompanhamento)
	if err != nil {
		a.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	a.JSON(http.StatusOK, response.SuccessResponse{
		Message: "Acompanhamento criado com sucesso",
	})
}

// BuscarPedido godoc
// @Summary Busca um pedido
// @Description Busca um pedido
// @Tags acompanhamento
// @Router /acompanhamento/{ID} [get]
// @Accept  json
// @Produce  json
// @Param ID path string true "ID do pedido"
// @Success 200 {object} presenters.PedidoDTO
// @Failure 400 {object} response.ErrorResponse
func (ah *AcompanhamentoHandler) BuscarPedido(p *gin.Context) {
	ID := p.Param("ID")
	fmt.Print(p.Params)
	pedido, err := ah.AcompanhamentoUseCase.BuscarPedido(p, ID)
	if err != nil {
		p.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	pedidoDTO := presenters.NewPedidoDTO(&pedido)
	p.JSON(http.StatusOK, pedidoDTO)
}

// AdicionarPedido godoc
// @Summary Adiciona um pedido ao acompanhamento
// @Description Adiciona um pedido existente ao acompanhamento de pedidos
// @Tags acompanhamento
// @Router /acompanhamento/{IDAcompanhamento}/{IDPedido} [post]
// @Accept json
// @Produce json
// @Param IDAcompanhamento path string true "ID do acompanhamento"
// @Param IDPedido path string true "ID do pedido"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse "Pedido ou acompanhamento não encontrado"
// @Failure 500 {object} response.ErrorResponse "Erro interno"
func (ah *AcompanhamentoHandler) AdicionarPedido(a *gin.Context) {
	idAcompanhamento := a.Param("IDAcompanhamento")
	idPedido := a.Param("IDPedido")

	pedido := &entities.Pedido{
		Identificacao: idPedido,
	}
	err := ah.AcompanhamentoUseCase.AdicionarPedido(a, idAcompanhamento, pedido)
	if err != nil {
		a.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	a.JSON(http.StatusOK, response.SuccessResponse{
		Message: "Pedido adicionado ao acompanhamento com sucesso",
	})
}

// BuscarAcompanhamento godoc
// @Summary Busca um acompanhamento
// @Description Busca um acompanhamento pelo ID
// @Tags acompanhamento
// @Router /acompanhamento/{ID} [get]
// @Accept json
// @Produce json
// @Param ID path string true "ID do acompanhamento"
// @Success 200 {object} presenters.AcompanhamentoDTO
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse "Acompanhamento não encontrado"
func (ah *AcompanhamentoHandler) BuscarAcompanhamento(c *gin.Context) {
	ID := c.Param("ID")
	acompanhamento, err := ah.AcompanhamentoUseCase.BuscarAcompanhamento(c, ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	acompanhamentoDTO := presenters.NewAcompanhamentoDTO(acompanhamento)
	c.JSON(http.StatusOK, acompanhamentoDTO)
}

// AtualizarStatusPedido godoc
// @Summary Atualiza o status de um pedido
// @Description Atualiza o status de um pedido no acompanhamento
// @Tags acompanhamento
// @Router /acompanhamento/{IDAcompanhamento}/pedido/{IDPedido}/status [put]
// @Accept json
// @Produce json
// @Param IDAcompanhamento path string true "ID do acompanhamento"
// @Param IDPedido path string true "ID do pedido"
// @Param status body StatusUpdateRequest true "Novo status do pedido"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse "Pedido ou acompanhamento não encontrado"
func (ah *AcompanhamentoHandler) AtualizarStatusPedido(a *gin.Context) {
	idAcompanhamento := a.Param("IDAcompanhamento")
	idPedido := a.Param("IDPedido")

	var statusUpdate StatusUpdateRequest
	if err := a.ShouldBindJSON(&statusUpdate); err != nil {
		a.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	err := ah.AcompanhamentoUseCase.AtualizarStatusPedido(a, idAcompanhamento, idPedido, entities.StatusPedido(statusUpdate.Status))
	if err != nil {
		a.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	a.JSON(http.StatusOK, response.SuccessResponse{
		Message: "Status do pedido atualizado com sucesso",
	})
}

type StatusUpdateRequest struct {
	Status string `json:"status" example:"Em preparação"`
}
