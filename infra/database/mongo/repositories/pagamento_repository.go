package repositories

import (
	"context"
	"fmt"

	"lanchonete/internal/domain/entities"
	"lanchonete/internal/domain/repository"
	"lanchonete/infra/database/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

type pagamentoMongoRepository struct {
	database   mongo.Database
	collection string
}

func NewPagamentoMongoRepository(db mongo.Database, collection string) repository.PagamentoRepository {
	return &pagamentoMongoRepository{
		database:   db,
		collection: collection,
	}
}

func (pr *pagamentoMongoRepository) EnviarPagamento(c context.Context, pagamento *entities.Pagamento) error {
	collection := pr.database.Collection(pr.collection)
	_, err := collection.InsertOne(c, pagamento)
	return err
}

func (pr *pagamentoMongoRepository) ConfirmarPagamento(c context.Context, pagamento *entities.Pagamento) error {
	collection := pr.database.Collection(pr.collection)

	filter := bson.M{"identificacao": pagamento.IdPagamento}
	update := bson.M{
		"$set": bson.M{
			"status":            pagamento.Status,
		},
	}

	result, err := collection.UpdateOne(c, filter, update)
	if err != nil {
		return fmt.Errorf("erro ao confirmar pagamento: %v", err)
	}

	if result.ModifiedCount == 0 {
		return fmt.Errorf("pagamento não encontrado")
	}

	return nil
} 