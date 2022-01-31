package config

import (
	"context"
	"fmt"
	"log"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

const (
	topic          = "guru"
	partition      = 0
	brokerAddress1 = "127.0.0.1:9092"
)

func CreateTopicKafka() {
	_, err := kafka.DialLeader(
		context.Background(),
		"tcp",
		brokerAddress1,
		topic, partition)

	if err != nil {
		log.Fatalf("Cannot create a topic %v", err)
	}

	fmt.Println("Created topic on kafka :", topic)
}

func PublishMessageFromKafka() {
	conn, err := kafka.DialLeader(context.Background(),
		"tcp", brokerAddress1,
		topic,
		partition)

	if err != nil {
		log.Fatalf("failed to dial leader %v:", err)
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte("one!")},
		kafka.Message{Value: []byte("two!")},
		kafka.Message{Value: []byte("three!")},
	)

	if err != nil {
		log.Fatalf("failed to write messages %v:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatalf("failed to close writer %v:", err)
	}
}

func ConsumeMessageFromKafka() {
	conn, err := kafka.DialLeader(context.Background(),
		"tcp",
		brokerAddress1,
		topic,
		partition)

	if err != nil {
		log.Fatalf("failed to dial leader %v:", err)
	}

	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

	b := make([]byte, 10e3) // 10KB max per message
	for {
		n, err := batch.Read(b)
		if err != nil {
			break
		}
		fmt.Println(string(b[:n]))
	}

	if err := batch.Close(); err != nil {
		log.Fatalf("failed to close batch %v:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatalf("failed to close connection %v:", err)
	}
}
