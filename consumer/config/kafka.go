package config

import (
	"log"
	"os"
	"strings"

	"github.com/IBM/sarama"
)

var (
	brokers = os.Getenv("KAFKA_BROKER")
	version = sarama.DefaultVersion.String()
	group   = os.Getenv("GROUP")
	topics  = os.Getenv("TOPIC")
	verbose = false
	oldest  = false
)

func InitConsumerGroup() sarama.ConsumerGroup {

	config := sarama.NewConfig()
	if oldest {
		config.Consumer.Offsets.Initial = sarama.OffsetOldest
	}
	config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRoundRobin()}

	client, err := sarama.NewConsumerGroup(strings.Split(brokers, ","), group, config)
	if err != nil {
		log.Panicf("new client: %v", err)
	}

	return client
}
