package config

import (
	"os"

	"github.com/IBM/sarama"
)

var (
	KafkaAddr = os.Getenv("KAFKA_BROKER")
)

func ProducerKafka() (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	producer, err := sarama.NewSyncProducer([]string{KafkaAddr}, config)
	return producer, err
}
