package main

import (
	"fmt"

	"os"
	"os/signal"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/tr2vil/its_manager/lib/config"
	. "github.com/tr2vil/its_manager/lib/logger"
	"github.com/tr2vil/its_manager/lib/message"

	"reflect"
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
	var consumer *kafka.Consumer
	var run bool = true

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	LogInit("../config/proc_stream.yaml")
	Log.Info("[Start Proc Stream]")
	fmt.Println(reflect.TypeOf(Log))

	conf := config.ConfigReader("../config/proc_stream.yaml")

	consumer = message.GetConsumer(Log, []string{"127.0.0.1:9092"})
	defer message.CloseConsumer(consumer)

	consumer.SubscribeTopics([]string{conf.ProcConf.Process.Topic_name}, nil)

	for run {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false

		case events := <-consumer.Events():
			switch e := events.(type) {
			case kafka.AssignedPartitions:
				fmt.Fprintf(os.Stderr, "%% %v\n", e)
				consumer.Assign(e.Partitions)
			case kafka.RevokedPartitions:
				fmt.Fprintf(os.Stderr, "%% %v\n", e)
				consumer.Unassign()
			case *kafka.Message:
				fmt.Printf("%% Message on %s:\n%s\n", e.TopicPartition, string(e.Value))
				var producer *kafka.Producer
				producer = message.GetProducer(Log, []string{"127.0.0.1:9092"})
				message.SendMessage2Archive(Log, producer, string(e.Value))
				message.CloseProducer(producer)
			case kafka.PartitionEOF:
				fmt.Printf("%% Reached %v\n", e)
			case kafka.Error:
				fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
			}
		}
	}

	Log.Info("[End Proc Stream]")
}
