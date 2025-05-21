package bootstrap

import (
	"context"
	"fmt"
	"log"

	"lanchonete/internal/domain/repository"
	"lanchonete/infra/database/mongo"
	"lanchonete/infra/database/mongo/repositories"
)

func NewDatabase(ctx context.Context, env *Env) (mongo.Database, error) {
	// Construct MongoDB URI from environment variables
	host := env.DBHost
	port := env.DBPort
	dbName := env.DBName

	if host == "" || port == "" || dbName == "" {
		return nil, fmt.Errorf("variáveis de ambiente DB_HOST, DB_PORT ou DB_NAME não estão definidas")
	}

	// Build MongoDB URI
	var mongoURI string
	mongoURI = fmt.Sprintf("mongodb://%s:%s", host, port)
	
	client, err := mongo.NewClient(mongoURI)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar cliente MongoDB: %v", err)
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar ao MongoDB: %v", err)
	}

	err = client.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer ping no MongoDB: %v", err)
	}

	log.Printf("Conectado ao MongoDB em %s!", mongoURI)
	return client.Database(dbName), nil
}

func NewRepositories(db mongo.Database) (
	repository.AcompanhamentoRepository,
	repository.PedidoRepository,
	repository.ProdutoRepository,
	repository.ClienteRepository,
	repository.PagamentoRepository,
) {
	return repositories.NewAcompanhamentoMongoRepository(db, "acompanhamento"),
		repositories.NewPedidoMongoRepository(db, "pedido"),
		repositories.NewProdutoMongoRepository(db, "produto"),
		repositories.NewClienteMongoRepository(db, "cliente"),
		repositories.NewPagamentoMongoRepository(db, "pagamento")
}

func CloseMongoDBConnection(client mongo.Client) {
	if client == nil {
		return
	}

	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to MongoDB closed.")
}