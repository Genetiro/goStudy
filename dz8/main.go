package main

import (
	conf "conf/conff/dz8/conff"
)

func main() {
	conf.Flconf(port int, db_url, jaeger_url, sentry_url, kafka_broker, some_app_id, some_app_key string)
}
