package main

import (
	"golang-kafka-clients/config"
)

func main() {
	// config.CreateTopicKafka()
	config.PublishMessageFromKafka()
	// config.ConsumeMessageFromKafka()
}
