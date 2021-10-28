package main

import (
	"fmt"

	"github.com/tr2vil/its_manager/lib"
)

func main() {
	common_conf := lib.ConfigReader("../config/common.yaml")
	fmt.Println(common_conf.Kafka.Hosts)
	proc_conf := lib.ConfigReader("../config/proc_archive.yaml")
	fmt.Println(proc_conf.Process.Name)
}
