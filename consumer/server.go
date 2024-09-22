package main

// SIGUSR1 toggle the pause/resume consumption
import (
	"context"
	"errors"
	_ "fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"savedb/config"
	"savedb/database"
	"savedb/skill"
	"strings"
	"sync"
	"syscall"

	"github.com/IBM/sarama"
)

// Sarama configuration options
var (
	brokers = os.Getenv("KAFKA_BROKER")
	version = sarama.DefaultVersion.String()
	group   = os.Getenv("GROUP")
	topics  = os.Getenv("TOPIC")
	verbose = false
	oldest  = false
)

func init() {
	if len(brokers) == 0 {
		panic("no Kafka bootstrap brokers defined, please set the -brokers flag")
	}
	if len(topics) == 0 {
		panic("no topics given to be consumed, please set the -topics flag")
	}
	if len(group) == 0 {
		panic("no Kafka consumer group defined, please set the -group flag")
	}
}
func main() {

	db := database.ConnectDB()
	defer db.Close()

	skillRepo := skill.NewSkillRepo(db)
	skillEventHandler := skill.NewSkillEventHandler(skillRepo)
	skillConsumer := skill.NewConsumerGroup(skillEventHandler)

	client := config.InitConsumerGroup()
	defer func() {
		if err := client.Close(); err != nil {
			log.Panicf("closing client: %v", err)
		}
	}()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	ctx, gracefully := context.WithCancel(context.Background())
	go func() {
		defer wg.Done()
		for {
			if err := client.Consume(ctx, strings.Split(topics, ","), skillConsumer); err != nil {
				if errors.Is(err, sarama.ErrClosedConsumerGroup) {
					return
				}
				log.Panicf("consume: %v", err)
			}
			// check if context was cancelled, signaling that the consumer should stop
			if err := ctx.Err(); err != nil {
				if errors.Is(err, context.Canceled) {
					slog.Info("the consumer context has cancelled for gracefully shutting down")
					return
				}
				slog.Error(ctx.Err().Error())
				return
			}
			slog.Info("rebalancing...")
			skillConsumer.NewReady()
		}
	}()
	<-skillConsumer.Ready()
	slog.Info("consumer up and running...")
	sigCtx, unregistered := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer unregistered()
keepRunning:
	for {
		select {
		case <-ctx.Done():
			slog.Info("terminating: consumer context cancel")
			break keepRunning
		case <-sigCtx.Done():
			slog.Info("terminating: via signal")
			unregistered()
			break keepRunning
		}
	}
	gracefully()
	wg.Wait() // waiting for gracefully consumer stopping
	if err := client.Close(); err != nil {
		log.Panicf("closing client: %v", err)
	}
}
