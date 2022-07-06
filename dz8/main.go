package main

import (
	conf "conf/conff/dz8/conff"
	"fmt"
)

func main() {
	fmt.Println(conf.GetEnv(DB_URL))
}
