# Kafka Exercise

This project simply sets up a local kafka cluster using Docker, a producer, and a consumer.
The producers and consumers constantly run at 2 messages per second until killed.

Based on: https://developer.confluent.io/get-started/go/ and adapted to use Shopify's 
Sarama Kafka client to avoid using CGO.

## Prerequisites

- docker daemon
- golang 1.20

## Running and building

### Go binaries

### Kafka and Docker setup

Build the local docker containers first and then run docker compose:

```
docker build --file producer/Dockerfile --tag kafka-producer .
docker build --file consumer/Dockerfile --tag kafka-consumer .
docker compose up -d
```

Look up the docker container IDs and then follow the log outputs, e.g.:

```
~/p/g/kafka-exercise ❯❯❯ docker ps -a | head -5
CONTAINER ID   IMAGE                              COMMAND                  CREATED             STATUS                         PORTS                                       NAMES
02420c66737d   confluentinc/cp-kafka:latest       "/etc/confluent/dock…"   12 seconds ago      Up 10 seconds                  0.0.0.0:9092->9092/tcp, :::9092->9092/tcp   broker
4f618a9d5ed9   confluentinc/cp-zookeeper:latest   "/etc/confluent/dock…"   12 seconds ago      Up 10 seconds                  2181/tcp, 2888/tcp, 3888/tcp                zookeeper
5c7d507a231c   kafka-consumer:latest              "./out/consumer"         12 seconds ago      Up 10 seconds                                                              consumer
615bfcd8680a   kafka-producer:latest              "./out/producer"         12 seconds ago      Up 10 seconds                                                              producer
~/p/g/kafka-exercise ❯❯❯ docker logs --follow 5c7d507a231c
2023/04/13 15:40:41 Consumed message offset 211: 'awalther': 'batteries'
2023/04/13 15:40:41 Consumed message offset 212: 'htanaka': 'alarm clock'
2023/04/13 15:40:42 Consumed message offset 213: 'htanaka': 'book'
2023/04/13 15:40:42 Consumed message offset 214: 'eabara': 'book'
[. . .]
```

In a separate terminal follow the producer to see the correlated messages being produced.

## TODO

- Get network configuration from file or environment variables
- Produce and consume >1 topics
- Use >1 Partition
