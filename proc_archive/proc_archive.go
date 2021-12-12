package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/tr2vil/its_manager/lib/config"
	. "github.com/tr2vil/its_manager/lib/logger"
	"github.com/tr2vil/its_manager/lib/message"
)

func main() {
	var consumer *kafka.Consumer
	LogInit("../config/proc_archive.yaml")
	Log.Info("[Start Proc Archive]")

	conf := config.ConfigReader("../config/proc_archive.yaml")
	fmt.Println(conf.ProcConf.Process.Name)

	consumer = message.GetConsumer([]string{"127.0.0.1:9092"})
	defer message.CloseConsumer(consumer)
	fmt.Println(consumer)

	Log.Info("[End Proc Archive]")
}
