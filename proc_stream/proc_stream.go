package main

import (
	"fmt"

	"github.com/tr2vil/its_manager/lib"
	. "github.com/tr2vil/its_manager/lib"
	//"github.com/tr2vil/its_manager/lib"
)

/*
func main() {
	LogInit("../config/proc_stream.yaml")
	conf := ConfigReader("../config/proc_stream.yaml")
	Log.Info("[Start Proc Stream]")

	ch := make(chan string, 10)
	quit := make(chan bool)
	go FileReadLine(conf.ProcConf.Process.Data_path, ch, quit)
	for {
		select {
		case res := <-ch:

			fmt.Println(res)
			WriteFile(conf.ProcConf.Process.Target_file, string(res+"\n"))
		}
	}
	fmt.Println("End for Loop")
}
*/

func main() {
	LogInit("../config/proc_stream.yaml")

	Log.Info("[Start Proc Stream]")

	conf := lib.ConfigReader("../config/proc_stream.yaml")

	consumer := lib.GetConsumer([]string{"127.0.0.1:9092"}, conf.ProcConf.Process.Topic_name)
	defer lib.CloseConsumer(consumer)
	fmt.Println(consumer)

	fmt.Println("End for Loop")
}
