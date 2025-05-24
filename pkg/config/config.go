package config

import "os"

// GetEnv retorna o valor de uma variável de ambiente ou um valor padrão se a variável não estiver definida.
// Se a variável de ambiente não estiver definida, retorna o valor padrão fornecido.
// Se a variável de ambiente estiver definida, retorna o valor dela.
// Exemplo de uso:
// value := GetEnv("MY_ENV_VAR", "default_value")
// Se MY_ENV_VAR estiver definido, value será o valor dela.
// Caso contrário, value será "default_value".
func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
