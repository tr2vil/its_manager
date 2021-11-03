package main

import (
	"fmt"

	"github.com/tr2vil/its_manager/lib"
)

func main() {
	conf := lib.ConfigReader("../config/proc_stream.yaml")

	ch := make(chan string, 10)
	quit := make(chan bool)
	go lib.FileReadLine(conf.ProcConf.Process.Data_path, ch, quit)
	for {
		select {
		case res := <-ch:

			fmt.Println(res)
			lib.WriteFile(conf.ProcConf.Process.Target_file, string(res+"\n"))
		}
	}
	fmt.Println("End for Loop")
}
