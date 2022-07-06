package conf

import (
	"os"
	"strconv"
)

type Flconf struct {
	port         int
	db_url       string
	jaeger_url   string
	sentry_url   string
	kafka_broker string
	some_app_id  string
	some_app_key string
}

func New() *Flconf {
	return &Flconf{

		port:         GetEnvAsInt("PORT", 8080),
		db_url:       GetEnv("DB_URL", "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable"),
		jaeger_url:   GetEnv("JAEGER_URL", "http://jaeger:16686"),
		sentry_url:   GetEnv("SENTRY_URL", "http://sentry:9000"),
		kafka_broker: GetEnv("KAFKA_BROKER", "kafka:9092"),
		some_app_id:  GetEnv("SOME_APP_ID", "testid"),
		some_app_key: GetEnv("SOME_APP_KEY", "testkey"),
	}
}
func GetEnvAsInt(name string, defaultVal int) int {
	valueStr := GetEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
}
func GetEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
