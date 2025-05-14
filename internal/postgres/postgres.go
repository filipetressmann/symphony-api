package postgres

import (
	"context"
	"fmt"
	"log"
	"symphony-api/pkg/config"

	"github.com/jackc/pgx/v5"
)

// InitPostgres inicializa a conexão com o banco de dados PostgreSQL e retorna o cliente.
// O cliente pode ser usado para executar consultas e interagir com o banco de dados.
// O URL de conexão é construído a partir das variáveis de ambiente definidas.
// As variáveis de ambiente esperadas são:
// POSTGRES_USER: Nome de usuário do PostgreSQL (padrão: "postgres")
// POSTGRES_PASSWORD: Senha do PostgreSQL (padrão: "password")
// POSTGRES_DB: Nome do banco de dados (padrão: "symphony")
// POSTGRES_HOST: Endereço do host do PostgreSQL (padrão: "localhost")
// POSTGRES_PORT: Porta do PostgreSQL (padrão: "5432")
// Se a conexão falhar, o programa será encerrado com um log de erro.
func InitPostgres() *pgx.Conn {
	user := config.GetEnv("POSTGRES_USER", "postgres")
	password := config.GetEnv("POSTGRES_PASSWORD", "password")
	dbName := config.GetEnv("POSTGRES_DB", "symphony")
	host := config.GetEnv("POSTGRES_HOST", "localhost")
	port := config.GetEnv("POSTGRES_PORT", "5432")

	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, dbName)

	client, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		log.Fatalf("Não foi possível conectar ao PostgreSQL: %v", err)
	}

	log.Println("Conectado ao PostgreSQL!")
	return client
}
