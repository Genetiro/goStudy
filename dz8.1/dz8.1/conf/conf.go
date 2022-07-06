package conf

import (
	"fmt"
	"net/url"
	"os"
)

type Urll url.URL

type Flconf struct {
	port         int    `defolt: 8080`
	db_url       Urll   `defolt: "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable"`
	jaeger_url   Urll   `defolt: "http://jaeger:16686"`
	sentry_url   Urll   `defolt: "http://sentry:9000"`
	kafka_broker string `defolt: "kafka:9092"`
	some_app_id  string `defolt: "testid"`
	some_app_key string `defolt: "testkey"`
}

/*func Conff() *Flconf {
	return &Flconf{

		port:         GetEnvAsInt("PORT", 8080),
		db_url:       GetEnvAsUrl("DB_URL", "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable"),
		jaeger_url:   GetEnvAsUrl("JAEGER_URL", "http://jaeger:16686"),
		sentry_url:   GetEnvAsUrl("SENTRY_URL", "http://sentry:9000"),
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
	return defaultVal
}
func GetEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func GetEnvAsUrl(name string, defaultVal Urll) (value Urll) {
	valStr := GetEnv(name, "")
	if valStr == "" {
		return defaultVal
	}
	return value
}*/

func (u *Urll) Decode(value string) error {
	*u, err := url.Parse(value)
	if err != nil {
		panic(err)
	}

	return nil
}
func Db() {
	var db Flconf
	fmt.Println(db)

}
