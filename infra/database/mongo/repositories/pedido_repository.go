package repositories

import (
	"context"
	"fmt"

	"lanchonete/internal/domain/entities"
	"lanchonete/internal/domain/repository"
	"lanchonete/infra/database/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

type pedidoMongoRepository struct {
	database   mongo.Database
	collection string
}

func NewPedidoMongoRepository(db mongo.Database, collection string) repository.PedidoRepository {
	return &pedidoMongoRepository{
		database:   db,
		collection: collection,
	}
}

func (pr *pedidoMongoRepository) CriarPedido(c context.Context, pedido *entities.Pedido) error {
	collection := pr.database.Collection(pr.collection)
	_, err := collection.InsertOne(c, pedido)
	return err
}

func (pr *pedidoMongoRepository) BuscarPedido(c context.Context, identificacao string) (*entities.Pedido, error) {
	collection := pr.database.Collection(pr.collection)

	filter := bson.M{
		"$or": []bson.M{
			{"identificacao": identificacao},
		},
	}

	var pedido entities.Pedido
	err := collection.FindOne(c, filter).Decode(&pedido)

	if err != nil {
		return nil, fmt.Errorf("pedido não encontrado: %v", err)
	}

	return &pedido, nil
}

func (pr *pedidoMongoRepository) AtualizarStatusPedido(c context.Context, Identificacao string, status string, UltimaAtualizacao string) error {
	collection := pr.database.Collection(pr.collection)

	filter := bson.M{"identificacao": Identificacao}
	update := bson.M{
		"$set": bson.M{
			"status":            status,
			"ultimaatualizacao": UltimaAtualizacao,
		},
	}

	result, err := collection.UpdateOne(c, filter, update)
	if err != nil {
		return fmt.Errorf("erro ao atualizar status do pedido: %v", err)
	}

	if result.ModifiedCount == 0 {
		return fmt.Errorf("pedido não encontrado")
	}

	return nil
}

func (pr *pedidoMongoRepository) ListarTodosOsPedidos(c context.Context) ([]*entities.Pedido, error) {
	collection := pr.database.Collection(pr.collection)

	cursor, err := collection.Find(c, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar pedidos: %v", err)
	}
	defer cursor.Close(c)

	var pedidos []*entities.Pedido
	if err = cursor.All(c, &pedidos); err != nil {
		return nil, fmt.Errorf("erro ao decodificar pedidos: %v", err)
	}

	return pedidos, nil
} 