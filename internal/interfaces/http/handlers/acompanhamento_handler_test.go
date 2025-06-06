package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"lanchonete/internal/domain/entities"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// --- Mock UseCases ---

type MockAcompanhamentoUseCase struct{ mock.Mock }

// Implements: AdicionarPedido(c context.Context, idAcompanhamento int, pedido *entities.Pedido) error
func (m *MockAcompanhamentoUseCase) AdicionarPedido(c context.Context, idAcompanhamento int, pedidoID int) error {
	args := m.Called(c, idAcompanhamento, pedidoID)
	return args.Error(0)
}

// Implements: BuscarAcompanhamento(c context.Context, id int) (*entities.AcompanhamentoPedido, error)
func (m *MockAcompanhamentoUseCase) BuscarAcompanhamento(c context.Context, id int) (*entities.AcompanhamentoPedido, error) {
	args := m.Called(c, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.AcompanhamentoPedido), args.Error(1)
}

// Implements: AtualizarStatusPedido(c context.Context, idAcompanhamento int, status entities.StatusPedido) error
func (m *MockAcompanhamentoUseCase) AtualizarStatusPedido(c context.Context, idAcompanhamento int, status entities.StatusPedido) error {
	args := m.Called(c, idAcompanhamento, status)
	return args.Error(0)
}

// Implements: CriarAcompanhamento(c context.Context, acompanhamento *entities.AcompanhamentoPedido) error
func (m *MockAcompanhamentoUseCase) CriarAcompanhamento(c context.Context) (int, error) {
	return 0, nil
}

// Implements: BuscarPedidos(c context.Context, idAcompanhamento int) ([]entities.Pedido, error)
func (m *MockAcompanhamentoUseCase) BuscarPedidos(c context.Context, idAcompanhamento int) ([]entities.Pedido, error) {
	args := m.Called(c, idAcompanhamento)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entities.Pedido), args.Error(1)
}

// --- Setup Handler ---

func setupAcompanhamentoHandlerWithMocks() (*AcompanhamentoHandler, *MockAcompanhamentoUseCase, *MockPedidoAtualizarStatusUseCase) {
	mockAcompanhamento := new(MockAcompanhamentoUseCase)
	mockPedidoAtualizar := new(MockPedidoAtualizarStatusUseCase)

	handler := &AcompanhamentoHandler{
		AcompanhamentoUseCase:        mockAcompanhamento,
		PedidoAtualizarStatusUseCase: mockPedidoAtualizar,
	}
	return handler, mockAcompanhamento, mockPedidoAtualizar
}

func TestAcompanhamentoHandler_AdicionarPedido(t *testing.T) {
	gin.SetMode(gin.TestMode)
	handler, mockAcompanhamento, _ := setupAcompanhamentoHandlerWithMocks()

	// Use integer IDs as expected by the mock
	mockAcompanhamento.On("AdicionarPedido", mock.Anything, 1, 2).Return(nil)

	req, _ := http.NewRequest(http.MethodPost, "/acompanhamento/1/2", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{
		{Key: "IDAcompanhamento", Value: "1"},
		{Key: "IDPedido", Value: "2"},
	}
	c.Request = req

	handler.AdicionarPedido(c)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Pedido adicionado ao acompanhamento com sucesso")
}

func TestAcompanhamentoHandler_BuscarAcompanhamento(t *testing.T) {
	gin.SetMode(gin.TestMode)
	handler, mockAcompanhamento, _ := setupAcompanhamentoHandlerWithMocks()

	acomp := &entities.AcompanhamentoPedido{ID: 1}
	mockAcompanhamento.On("BuscarAcompanhamento", mock.Anything, 1).Return(acomp, nil)

	req, _ := http.NewRequest(http.MethodGet, "/acompanhamento/1", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "ID", Value: "1"}}
	c.Request = req

	handler.BuscarAcompanhamento(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAcompanhamentoHandler_AtualizarStatusPedido(t *testing.T) {
	gin.SetMode(gin.TestMode)
	handler, mockAcompanhamento, _ := setupAcompanhamentoHandlerWithMocks()

	statusReq := StatusUpdateRequest{Status: "Finalizado"}
	mockAcompanhamento.On("AtualizarStatusPedido", mock.Anything, 1, entities.StatusPedido("Finalizado")).Return(nil)
	mockAcompanhamento.On("BuscarPedidos", mock.Anything, 2).Return([]entities.Pedido{}, nil)

	body, _ := json.Marshal(statusReq)
	req, _ := http.NewRequest(http.MethodPut, "/acompanhamento/1/pedido/2/status", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{
		{Key: "IDAcompanhamento", Value: "1"},
		{Key: "IDPedido", Value: "2"},
	}
	c.Request = req

	handler.AtualizarStatusPedido(c)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Status do pedido atualizado com sucesso")
}
