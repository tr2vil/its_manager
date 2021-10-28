package main

import (
	"fmt"

	"github.com/tr2vil/its_manager/lib"
)

func main() {
	/*
		common_conf := lib.CommonConfigReader("../config/common.yaml")
		fmt.Println(common_conf.Kafka.Hosts)
		proc_conf := lib.ProcessConfigReader("../config/proc_stream.yaml")
		fmt.Println(proc_conf.Process.Script_path)
	*/
	conf := lib.ConfigReader("../config/proc_stream.yaml")
	fmt.Println(conf.CommonConf.Log.Level)
	fmt.Println(conf.CommonConf.Log.Format)
	fmt.Println(conf.ProcConf.Process.Name)
	fmt.Println(conf.ProcConf.Process.Topic_name)
}
