package repositories

import (
	"context"
	"fmt"
	"time"

	"lanchonete/internal/domain/entities"
	"lanchonete/internal/domain/repository"
	"lanchonete/infra/database/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type acompanhamentoMongoRepository struct {
	database          mongo.Database
	collection        string
	pedidosCollection string
}

func NewAcompanhamentoMongoRepository(db mongo.Database, collection string) repository.AcompanhamentoRepository {
	return &acompanhamentoMongoRepository{
		database:          db,
		collection:        collection,
		pedidosCollection: "pedido",
	}
}

func (ar *acompanhamentoMongoRepository) CriarAcompanhamento(c context.Context, acompanhamento *entities.AcompanhamentoPedido) error {
	collection := ar.database.Collection(ar.collection)
	_, err := collection.InsertOne(c, acompanhamento)
	return err
}

func (ar *acompanhamentoMongoRepository) BuscarPedidos(c context.Context, ID string) (entities.Pedido, error) {
	collection := ar.database.Collection(ar.pedidosCollection)

	filter := bson.M{
		"$or": []bson.M{
			{"identificacao": ID},
		},
	}

	var pedido entities.Pedido
	err := collection.FindOne(c, filter).Decode(&pedido)

	if err != nil {
		return entities.Pedido{}, fmt.Errorf("pedido não encontrado: %v", err)
	}

	return pedido, nil
}

func (ar *acompanhamentoMongoRepository) AdicionarPedido(c context.Context, acompanhamento *entities.AcompanhamentoPedido, p *entities.Pedido) error {
	// Buscar o pedido
	pedido, err := ar.BuscarPedidos(c, p.Identificacao)
	if err != nil {
		return fmt.Errorf("erro ao buscar pedido: %v", err)
	}

	// Construir o filtro
	filter := bson.M{"id": acompanhamento.ID}

	// Adicionar o pedido à fila
	acompanhamento.Pedidos.Enfileirar(pedido)

	// Obter a lista de pedidos atualizada
	pedidosAtualizados := acompanhamento.Pedidos.Listar()

	// Atualizar o documento
	local, _ := time.LoadLocation("America/Sao_Paulo")
	agora := time.Now().In(local).Format("2006-01-02 15:04:05")
	update := bson.M{
		"$set": bson.M{
			"pedidos":           bson.M{"pedidos": pedidosAtualizados},
			"ultimaAtualizacao": agora,
		},
	}

	result, err := ar.database.Collection(ar.collection).UpdateOne(c, filter, update)
	if err != nil {
		return fmt.Errorf("erro ao atualizar acompanhamento: %v", err)
	}

	if result.ModifiedCount == 0 {
		return fmt.Errorf("nenhum documento foi atualizado")
	}

	return nil
}

func (ar *acompanhamentoMongoRepository) BuscarAcompanhamento(ctx context.Context, ID string) (*entities.AcompanhamentoPedido, error) {
	var acompanhamento *entities.AcompanhamentoPedido

	filter := bson.M{"id": ID}

	err := ar.database.Collection(ar.collection).FindOne(ctx, filter).Decode(&acompanhamento)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar acompanhamento: %v", err)
	}

	acompanhamento.Pedidos.EnsureMutex()
	return acompanhamento, nil
}

func (ar *acompanhamentoMongoRepository) AtualizarStatusPedido(ctx context.Context, acompanhamentoID string, identificacao string, novoStatus entities.StatusPedido) error {
	filter := bson.M{"id": acompanhamentoID}

	local, _ := time.LoadLocation("America/Sao_Paulo")
	agora := time.Now().In(local).Format("2006-01-02 15:04:05")
	update := bson.M{
		"$set": bson.M{
			"pedidos.pedidos.$[elem].status":            string(novoStatus),
			"pedidos.pedidos.$[elem].ultimaatualizacao": agora,
			"ultimaAtualizacao":                         agora,
		},
	}

	// Remover se finalizado
	if novoStatus == entities.Finalizado {
		removeUpdate := bson.M{
			"$pull": bson.M{
				"pedidos.pedidos": bson.M{"identificacao": identificacao},
			},
		}
		_, err := ar.database.Collection(ar.collection).UpdateOne(ctx, filter, removeUpdate)
		if err != nil {
			return fmt.Errorf("erro ao remover pedido finalizado: %v", err)
		}
	}

	arrayFilters := options.Update().SetArrayFilters(options.ArrayFilters{
		Filters: []interface{}{
			bson.M{"elem.identificacao": identificacao},
		},
	})

	result, err := ar.database.Collection(ar.collection).UpdateOne(ctx, filter, update, arrayFilters)
	if err != nil {
		return fmt.Errorf("erro ao atualizar pedido: %v", err)
	}

	if result.ModifiedCount == 0 {
		return fmt.Errorf("pedido %s não encontrado no acompanhamento %s", identificacao, acompanhamentoID)
	}

	return nil
} 