package lib

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func GetConsumer(hosts []string, topic string) *kafka.Consumer {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": hosts,
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	return consumer
}
