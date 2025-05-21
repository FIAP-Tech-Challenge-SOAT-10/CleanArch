package repositories

import (
	"context"
	"fmt"

	"lanchonete/internal/domain/entities"
	"lanchonete/internal/domain/repository"
	"lanchonete/infra/database/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

type produtoMongoRepository struct {
	database   mongo.Database
	collection string
}

func NewProdutoMongoRepository(db mongo.Database, collection string) repository.ProdutoRepository {
	return &produtoMongoRepository{
		database:   db,
		collection: collection,
	}
}

func (pr *produtoMongoRepository) AdicionarProduto(c context.Context, produto *entities.Produto) error {
	collection := pr.database.Collection(pr.collection)
	_, err := collection.InsertOne(c, produto)
	return err
}

func (pr *produtoMongoRepository) BuscarProdutoPorId(c context.Context, identificacao string) (*entities.Produto, error) {
	collection := pr.database.Collection(pr.collection)

	filter := bson.M{"identificacao": identificacao}

	var produto entities.Produto
	err := collection.FindOne(c, filter).Decode(&produto)

	if err != nil {
		return nil, fmt.Errorf("produto não encontrado: %v", err)
	}

	return &produto, nil
}

func (pr *produtoMongoRepository) ListarTodosOsProdutos(c context.Context) ([]*entities.Produto, error) {
	collection := pr.database.Collection(pr.collection)

	cursor, err := collection.Find(c, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar produtos: %v", err)
	}
	defer cursor.Close(c)

	var produtos []*entities.Produto
	if err = cursor.All(c, &produtos); err != nil {
		return nil, fmt.Errorf("erro ao decodificar produtos: %v", err)
	}

	return produtos, nil
}

func (pr *produtoMongoRepository) EditarProduto(c context.Context, produto *entities.Produto) error {
	collection := pr.database.Collection(pr.collection)

	filter := bson.M{"nome": produto.Nome}
	update := bson.M{"$set": produto}

	result, err := collection.UpdateOne(c, filter, update)
	if err != nil {
		return fmt.Errorf("erro ao atualizar produto: %v", err)
	}

	if result.ModifiedCount == 0 {
		return fmt.Errorf("produto não encontrado")
	}

	return nil
}

func (pr *produtoMongoRepository) RemoverProduto(c context.Context, identificacao string) error {
	collection := pr.database.Collection(pr.collection)

	filter := bson.M{"identificacao": identificacao}

	deletedCount, err := collection.DeleteOne(c, filter)
	if err != nil {
		return fmt.Errorf("erro ao remover produto: %v", err)
	}

	if deletedCount == 0 {
		return fmt.Errorf("produto não encontrado")
	}

	return nil
}

func (pr *produtoMongoRepository) ListarPorCategoria(c context.Context, categoria string) ([]*entities.Produto, error) {
	collection := pr.database.Collection(pr.collection)

	filter := bson.M{"categoria": categoria}

	cursor, err := collection.Find(c, filter)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar produtos por categoria: %v", err)
	}
	defer cursor.Close(c)

	var produtos []*entities.Produto
	if err = cursor.All(c, &produtos); err != nil {
		return nil, fmt.Errorf("erro ao decodificar produtos: %v", err)
	}

	return produtos, nil
} 