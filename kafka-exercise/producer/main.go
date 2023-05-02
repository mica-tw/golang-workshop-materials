package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/Shopify/sarama"
)

func main() {
	// wait for broker to be ready
	time.Sleep(10 * time.Second)

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	p, err := sarama.NewAsyncProducer([]string{"broker:9092"}, config)
	if err != nil {
		fmt.Printf("Failed to create producer: %s", err)
		os.Exit(1)
	}

	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// Go-routine to handle message delivery reports and
	// possibly other event types (errors, stats, etc)
	successes := 0
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for msg := range p.Successes() {
			successes++

			log.Printf("Sent message: %v\n", *msg)
		}
	}()

	myChan := make(chan int, 1)
	myChan <- 5
	close(myChan)
	<-myChan

	myChan <- 6

	producerErrors := 0
	wg.Add(1)
	go func() {
		defer wg.Done()
		for err := range p.Errors() {
			log.Println(err)
			producerErrors++
		}
	}()

	topic := "purchases"
	users := [...]string{"eabara", "jsmith", "sgarcia", "jbernard", "htanaka", "awalther"}
	items := [...]string{"book", "alarm clock", "t-shirts", "gift card", "batteries"}

	wg.Add(1)
	go func() {
		defer wg.Done()

		ticker := time.Tick(500 * time.Millisecond)

		for {
			select {
			case <-ticker:
				key := users[rand.Intn(len(users))]
				data := items[rand.Intn(len(items))]
				msg := sarama.ProducerMessage{
					Topic: topic,
					Key:   sarama.StringEncoder(key),
					Value: sarama.StringEncoder(data),
				}
				p.Input() <- &msg
			case <-signals:
				p.AsyncClose()
				return
			}
		}
	}()

	// Wait for all messages to be delivered
	wg.Wait()

	log.Printf("Successfully produced: %d; errors: %d\n", successes, producerErrors)
}
