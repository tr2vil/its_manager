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
		}
	}
	fmt.Println("End for Loop")
}
