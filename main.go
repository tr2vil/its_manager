package main

import (
	"fmt"

	//"github.com/takama/daemon"
	"github.com/tr2vil/its_manager/lib"
)

/*
func main() {
	service, err := daemon.New("name_1", "desc_1", daemon.SystemDaemon)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	status, err := service.Install()
	if err != nil {
		log.Fatal(status, "\nError: ", err)
	}

	fmt.Println(status)
}
*/

func main() {
	conf := lib.ConfigReader("./config/properties.yaml")
	fmt.Println(conf.Kafka.Hosts)
}
