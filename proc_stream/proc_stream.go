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
	fmt.Println("Start for Loop")
	for {
		select {
		case res := <-ch:
			fmt.Println("1-", res)
		}
	}
	fmt.Println("End for Loop")
}
