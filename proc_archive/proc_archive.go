package main

import (
	"fmt"

	"github.com/tr2vil/its_manager/lib"
)

func main() {
	conf := lib.ConfigReader("../config/properties.yaml")
	fmt.Println(conf.Kafka.Hosts)
	fmt.Println(conf.Process.Name)
}
