package repositories

import (
	"context"
	"fmt"

	"lanchonete/internal/domain/entities"
	"lanchonete/internal/domain/repository"
	"lanchonete/infra/database/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

type clienteMongoRepository struct {
	database   mongo.Database
	collection string
}

func NewClienteMongoRepository(db mongo.Database, collection string) repository.ClienteRepository {
	return &clienteMongoRepository{
		database:   db,
		collection: collection,
	}
}

func (cr *clienteMongoRepository) CriarCliente(c context.Context, cliente *entities.Cliente) error {
	collection := cr.database.Collection(cr.collection)
	_, err := collection.InsertOne(c, cliente)
	return err
}

func (cr *clienteMongoRepository) BuscarCliente(c context.Context, CPF string) (*entities.Cliente, error) {
	collection := cr.database.Collection(cr.collection)

	filter := bson.M{"cpf": CPF}

	var cliente entities.Cliente
	err := collection.FindOne(c, filter).Decode(&cliente)

	if err != nil {
		return nil, fmt.Errorf("cliente não encontrado: %v", err)
	}

	return &cliente, nil
} 