package mongo

import (
	"context"
	"fmt"
	"log"
	"symphony-api/pkg/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InitMongo inicializa a conexão com o banco de dados MongoDB e retorna o cliente.
// O cliente pode ser usado para executar consultas e interagir com o banco de dados.
// O URL de conexão é construído a partir das variáveis de ambiente definidas.
// As variáveis de ambiente esperadas são:
// MONGO_INITDB_ROOT_USERNAME: Nome de usuário do MongoDB (padrão: "root")
// MONGO_INITDB_ROOT_PASSWORD: Senha do MongoDB (padrão: "rootpassword")
func InitMongo() *mongo.Client {
	username := config.GetEnv("MONGO_INITDB_ROOT_USERNAME", "root")
	password := config.GetEnv("MONGO_INITDB_ROOT_PASSWORD", "rootpassword")
	mongoURI := fmt.Sprintf("mongodb://%s:%s@localhost:27017", username, password)

	clientOptions := options.Client().ApplyURI(mongoURI)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("Falha a se conectar ao MongoDB: %v", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("Falha ao pingar o MongoDB: %v", err)
	}

	log.Println("Conectado ao MongoDB!")
	return client
}
