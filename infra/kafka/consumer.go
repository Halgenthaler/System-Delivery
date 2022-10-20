package kafka

import (
	"fmt"
	"log"
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type kafkaConsumer struct {
	MsgChan chan *ckafka.Message
}

func (k *kafkaConsumer) Consume() {
	configmap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KafkaBootstrapServer"),
		"group.id":          os.Getenv("KafkaConsumerGroupId"),
	}

	c, err := ckafka.NewConsumer(configmap)
	if err != nil {
		log.Fatalf("error consume kafka: " + err.Error())
	}

	topics := []string{os.Getenv("KafkaReadTopic")}
	c.SubscribeTopics(topics, nil)
	fmt.Println("Consumer has been started")

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			k.MsgChan <- msg
		}
	}

}
