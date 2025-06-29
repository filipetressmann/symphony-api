package postgres

import (
	"context"
	"fmt"
	"log"
	"strings"
	"symphony-api/pkg/config"

	"github.com/jackc/pgx/v5"
)

type PostgreConnection interface {
	Put(data map[string]any, tableName string) error
	PutReturningId(data map[string]any, tableName string, idName string) (any, error)
	Get(constraints map[string]any, tableName string) ([]map[string]any, error)
	GetChatWithLimit(chat_id int32, limit int32, tableName string) ([]map[string]any, error)
}

type PostgreConnectionImpl struct {
	*pgx.Conn
}

func NewPostgreConnection() PostgreConnection {
	user := config.GetEnv("POSTGRES_USER", "user")
	password := config.GetEnv("POSTGRES_PASSWORD", "password")
	dbName := config.GetEnv("POSTGRES_DB", "symphony")
	host := config.GetEnv("POSTGRES_HOST", "postgres")
	port := config.GetEnv("POSTGRES_PORT", "5432")

	dbUrl := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbName)
	
	client, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		log.Fatal("Failed to connect to postgres: ", err)
	}

	log.Println("Successfully connected to postgres!")
	return &PostgreConnectionImpl{
		Conn: client,
	}
}

func (conn *PostgreConnectionImpl) Put(data map[string]any, tableName string) error {
	insertStatement, args := getInsertStament(data, tableName, nil)
	log.Printf("Executing insert statement at Postgres: %s", insertStatement)
	_, err := conn.Exec(
		context.Background(),
		insertStatement,
		args...,
	)

	return err
}

func (conn *PostgreConnectionImpl) PutReturningId(data map[string]any, tableName string, idName string) (any, error) {
	var id any

	insertStatement, args := getInsertStament(data, tableName, &idName)
	log.Printf("Executing insert statement at Postgres: %s", insertStatement)
	err := conn.QueryRow(
		context.Background(),
		insertStatement,
		args...,
	).Scan(&id)

	return id, err
}

func getInsertStament(data map[string]any, tableName string, idName *string) (string, []any) {
	if len(data) == 0 {
        if idName != nil {
            return fmt.Sprintf("INSERT INTO %s DEFAULT VALUES RETURNING %s", tableName, *idName), nil
        }
        return fmt.Sprintf("INSERT INTO %s DEFAULT VALUES", tableName), nil
    }

	keys := make([]string, 0, len(data))
	values := make([]any, 0, len(data))
	placeholders := make([]string, 0, len(data))

	baseInsert := "INSERT INTO %s (%s) VALUES (%s)"

	if idName != nil {
		baseInsert = fmt.Sprintf("%s RETURNING %s", baseInsert, *idName)
	}

	index := 1
	for k, v := range data {
		keys = append(keys, k)
		values = append(values, v)
		placeholders = append(placeholders, fmt.Sprintf("$%d", index))
		index += 1
	}

	return fmt.Sprintf(
		baseInsert, 
		tableName, 
		joinComma(keys), 
		joinComma(placeholders),
	), values
}

func (conn *PostgreConnectionImpl) Get(constraints map[string]any, tableName string) ([]map[string]any, error) {
	sql, args := getSelectWthConstraintsQuery(constraints, tableName)

	rows, err := conn.Query(
		context.Background(),
		sql,
		args...,
	)

	if err != nil {
		return nil, err
	}

	result, err := rowsToMaps(rows)

	return result, err
}

func rowsToMaps(rows pgx.Rows) ([]map[string]any, error) {
	defer rows.Close()

	columns := getRowColumns(rows)

	var results []map[string]interface{}

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return nil, err
		}

		rowMap := make(map[string]interface{})
		for i, col := range columns {
			rowMap[col] = values[i]
		}

		results = append(results, rowMap)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return results, nil
}

func getRowColumns(rows pgx.Rows) []string {
	fieldDescriptions := rows.FieldDescriptions()
	columns := make([]string, len(fieldDescriptions))
	for i, fd := range fieldDescriptions {
		columns[i] = string(fd.Name)
	}

	return columns
}

func getSelectWthConstraintsQuery(constraints map[string]any, tableName string) (string, []any) {
	values := make([]any, 0, len(constraints))
	constraintList := make([]string, 0, len(constraints))

	index := 1
	for k, v := range constraints {
		constraintList = append(constraintList, fmt.Sprintf("%s = $%d", k, index))
		values = append(values, v)
		index += 1
	}

	return fmt.Sprintf(
			"SELECT * FROM %s WHERE %s", 
			tableName, 
			joinComma(constraintList),
		), values
}

func joinComma(values []string) string {
	return strings.Join(values, ",")
}

func (conn *PostgreConnectionImpl) GetChatWithLimit(chat_id int32, limit int32, tableName string) ([]map[string]any, error) {
    var time_column string
    switch tableName {
    case "CHAT_MESSAGE":
        time_column = "sent_at"
    case "CHAT":
        time_column = "created_at"
    default:
        return nil, fmt.Errorf("unsupported table: %s", tableName)
    }

    sql := fmt.Sprintf("SELECT * FROM %s WHERE chat_id = $1 ORDER BY %s DESC LIMIT $2", tableName, time_column)
    rows, err := conn.Query(context.Background(), sql, chat_id, limit)
    if err != nil {
        return nil, err
    }
    return rowsToMaps(rows)
}