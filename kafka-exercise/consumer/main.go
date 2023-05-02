package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/Shopify/sarama"
)

func main() {
	time.Sleep(10 * time.Second)

	consumer, err := sarama.NewConsumer([]string{"broker:9092"}, sarama.NewConfig())
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	topic := "purchases"
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	consumed := 0
	notKilled := true

	for notKilled {
		// blocks before select until there is a msg in at least one chan
		select {
		case msg := <-partitionConsumer.Messages():
			log.Printf("Consumed message offset %d: '%s': '%s'\n", msg.Offset, string(msg.Key), string(msg.Value))
			consumed++
		case <-signals:
			notKilled = false
		}
	}

	log.Printf("Consumed: %d\n", consumed)
}
