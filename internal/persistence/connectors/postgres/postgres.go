package postgres

import (
	"context"
	"fmt"
	"log"
	"strings"
	"symphony-api/pkg/config"

	"github.com/jackc/pgx/v5"
)

// PostgreConnection defines the interface for interacting with a PostgreSQL database.
// It provides methods for inserting data into a table and retrieving data with constraints.
// The Put method inserts a map of data into a specified table and returns the ID of the inserted row.
// The Get method retrieves rows from a specified table based on given constraints.
type PostgreConnection interface {
	Put(data map[string]any, tableName string) (int64, error)
	Get(constraints map[string]any, tableName string) []map[string]any
}

// PostgreConnectionImpl is a concrete implementation of the PostgreConnection interface.
// It holds a pgx.Conn instance which is used to interact with the PostgreSQL database.
// The Put method constructs an SQL insert statement from the provided data map,
// executes it, and returns the ID of the newly inserted row.
// The Get method constructs an SQL select statement with the provided constraints,
// executes it, and returns the results as a slice of maps, where each map represents a row.
// The PostgreConnectionImpl struct is designed to be used with the pgx library for PostgreSQL.
// It is initialized with a connection to the PostgreSQL database, which is established using
// connection parameters read from environment variables.
type PostgreConnectionImpl struct {
	*pgx.Conn
}

// InitPostgres initializes a new PostgreSQL connection.
// It reads the connection parameters from environment variables,
// constructs the connection URL, and creates a PostgreSQL client.
// If the connection fails, it logs the error and exits the application.
// The PostgreConnectionImpl struct holds the PostgreSQL connection which can be used
// to interact with the PostgreSQL database.
// The connection parameters are:
// - POSTGRES_USER: The username for the PostgreSQL database (default: "user").
// - POSTGRES_PASSWORD: The password for the PostgreSQL database (default: "password").
// - POSTGRES_DB: The name of the PostgreSQL database (default: "symphony").
// - POSTGRES_HOST: The host of the PostgreSQL database (default: "postgres").
// - POSTGRES_PORT: The port of the PostgreSQL database (default: "5432").
// Returns a PostgreConnection instance.
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

// Put inserts a map of data into a specified table in the PostgreSQL database.
// It constructs an SQL insert statement from the provided data map,
// executes it, and returns the ID of the newly inserted row.
// The data map should contain key-value pairs where keys are column names
// and values are the corresponding values to be inserted.
// The tableName parameter specifies the name of the table where the data should be inserted.
// The method returns the ID of the inserted row and an error if the operation fails.
func (conn *PostgreConnectionImpl) Put(data map[string]any, tableName string) (int64, error) {
	var id int64

	insertStatement, args := getInsertStament(data, tableName)
	log.Printf("Executing insert statement at Postgres: %s", insertStatement)
	err := conn.QueryRow(
		context.Background(),
		insertStatement,
		args...,
	).Scan(&id)

	return id, err
}

// getInsertStament constructs an SQL insert statement from the provided data map.
// It returns the SQL statement as a string and a slice of values to be used as arguments in the query.
func getInsertStament(data map[string]any, tableName string) (string, []any) {
	keys := make([]string, 0, len(data))
	values := make([]any, 0, len(data))
	placeholders := make([]string, 0, len(data))

	index := 1
	for k, v := range data {
		keys = append(keys, k)
		values = append(values, v)
		placeholders = append(placeholders, fmt.Sprintf("$%d", index))
		index += 1
	}

	return fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s) RETURNING id", 
		tableName, 
		joinComma(keys), 
		joinComma(placeholders),
	), values
}

// Get retrieves rows from a specified table in the PostgreSQL database based on given constraints.
// It constructs an SQL select statement with the provided constraints,
// executes it, and returns the results as a slice of maps, where each map represents a row.
// The constraints map should contain key-value pairs where keys are column names
// and values are the corresponding values to filter the rows.
// The tableName parameter specifies the name of the table from which to retrieve the data.
// The method returns a slice of maps representing the rows that match the constraints.
func (conn *PostgreConnectionImpl) Get(constraints map[string]any, tableName string) []map[string]any {
	sql, args := getSelectWthConstraintsQuery(constraints, tableName)

	rows, err := conn.Query(
		context.Background(),
		sql,
		args,
	)
	
	if err != nil && err != pgx.ErrNoRows {
		log.Fatal("Query failed:", err)
	}

	result, err := rowsToMaps(rows)

	if err != nil {
		log.Fatal("Row conversion failed: ", err)
	}

	return result
}

// rowsToMaps converts pgx.Rows to a slice of maps, where each map represents a row.
// It retrieves the column names from the rows and maps each row's values to the corresponding column names.
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

// getRowColumns retrieves the column names from pgx.Rows.
// It returns a slice of strings representing the column names.
func getRowColumns(rows pgx.Rows) []string {
	fieldDescriptions := rows.FieldDescriptions()
	columns := make([]string, len(fieldDescriptions))
	for i, fd := range fieldDescriptions {
		columns[i] = string(fd.Name)
	}

	return columns
}

// getSelectWthConstraintsQuery constructs an SQL select statement with the provided constraints.
// It returns the SQL statement as a string and a slice of values to be used as arguments in the query.
func getSelectWthConstraintsQuery(constraints map[string]any, tableName string) (string, []any) {
	values := make([]any, 0, len(constraints))
	constraintList := make([]string, 0, len(constraints))

	index := 1
	for k, v := range constraints {
		constraintList = append(constraintList, fmt.Sprintf("%s = %d", k, index))
		values = append(values, v)
		index += 1
	}

	return fmt.Sprintf(
			"SELECT * FROM %s WHERE %s", 
			tableName, 
			joinComma(constraintList),
		), values
}

// joinComma joins a slice of strings with commas.
// It returns a single string with the elements of the slice separated by commas.
func joinComma(values []string) string {
	return strings.Join(values, ",")
}