package conf

import (
	"flag"
	"fmt"
	"net/url"
)

type flconf struct {
	port         int
	db_url       string
	jaeger_url   string
	sentry_url   string
	kafka_broker string
	some_app_id  string
	some_app_key string
}

func Configg() {
	conf := flconf{}
	flag.IntVar(&conf.port, "Port", 8080, "port")
	flag.StringVar(&conf.db_url, "DB_url", "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable", "DB_url")
	flag.StringVar(&conf.jaeger_url, "Jaeger_url", "http://jaeger:16686", "Jaeger_url")
	flag.StringVar(&conf.sentry_url, "Sentry_url", "http://sentry:9000", "Sentry_url")
	flag.StringVar(&conf.kafka_broker, "Kafka_broker", "kafka:9092", "Kafka_broker")
	flag.StringVar(&conf.some_app_id, "Some_app_id", "testid", "Some_app_id")
	flag.StringVar(&conf.some_app_key, "Some_app_key", "testkey", "Some_app_key")

	flag.Parse()
	flconf(conf.port, conf.db_url, conf.jaeger_url, conf.sentry_url, conf.kafka_broker, conf.some_app_id, conf.some_app_key)

}

func Flconf(port int, db_url, jaeger_url, sentry_url, kafka_broker, some_app_id, some_app_key string) {
	_, err := url.Parse(db_url)
	if err != nil {
		panic(err)
	}
	fmt.Println(port, db_url, jaeger_url, sentry_url, kafka_broker, some_app_id, some_app_key)
}
