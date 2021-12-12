package message

import (
	"strings"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sirupsen/logrus"
)

type kafkaData struct {
	Log    *logrus.Logger
	c      *kafka.Consumer
	p      *kafka.Producer
	broker []string
}

func GetConsumer(Log *logrus.Logger, broker []string) *kafka.Consumer {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":               strings.Join(broker, ","),
		"group.id":                        "its-manager",
		"session.timeout.ms":              6000,
		"go.events.channel.enable":        true,
		"go.application.rebalance.enable": true,
		"enable.partition.eof":            true,
		"auto.offset.reset":               "earliest",
	})
	if err != nil {
		Log.Error("[Error] Kafka NewConsumer Fail!")
		panic(err)
	}

	return consumer
}

func GetProducer(Log *logrus.Logger, broker []string) *kafka.Producer {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker[0]})
	if err != nil {
		Log.Error("[Error] Kafka NewProducer Fail!")
		panic(err)
	}

	Log.Debug("[Debug] Created Producer : %v", producer)

	return producer
}

func CloseConsumer(consumer *kafka.Consumer) {
	if consumer != nil {
		consumer.Close()
		consumer = nil
	}
}

func CloseProducer(producer *kafka.Producer) {
	if producer != nil {
		producer.Close()
		producer = nil
	}
}

func sendMessage(Log *logrus.Logger, producer *kafka.Producer, topic string, msg string) {
	producer.ProduceChannel() <- &kafka.Message{TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny}, Value: []byte(msg)}
}

func SendMessage2Archive(Log *logrus.Logger, producer *kafka.Producer, msg string) {
	var topic string = "its-archive"
	sendMessage(Log, producer, topic, msg)
}
