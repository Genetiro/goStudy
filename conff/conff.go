package conff

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
)

type Urll url.URL
type Conff struct {
	Port         int    `json: "port"`
	Db_url       string `json: "db_url string"`
	Jaeger_url   string `json: "jaeger_url"`
	Sentry_url   string `json: "sentry_url"`
	Kafka_broker string `json: "kafka_broker"`
	Some_app_id  string `json: "some_app_id"`
	Some_app_key string `json: "some_app_key"`
}

/*func (u *Urll) Decode(val string) error {

	u, err := url.Parse(val)

	if err != nil {
		panic(err)
	}
	return err
}*/

func GetConff() {
	data, err := os.ReadFile("conff.json")
	if err != nil {
		panic(err)
	}
	var cnfg Conff
	if err = json.Unmarshal(data, &cnfg); err != nil {
		panic(err)
	}
	fmt.Printf("% + v\n", cnfg)
}
