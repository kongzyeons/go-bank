package kafka

import (
	"log"

	"github.com/IBM/sarama"
)

func InitKafkaProducer() sarama.SyncProducer {
	servers := []string{"localhost:9092"}
	producer, err := sarama.NewSyncProducer(servers, nil)

	if err != nil {
		log.Fatalf("Failed to connect to kafka: %v", err)
	}
	return producer

}

func InitKafkaConsumer() sarama.ConsumerGroup {
	servers := []string{"localhost:9092"}
	groupID := "accountConsumer"
	consumer, err := sarama.NewConsumerGroup(servers, groupID, nil)
	if err != nil {
		log.Fatalf("Failed to connect to kafka: %v", err)
	}
	return consumer
}
