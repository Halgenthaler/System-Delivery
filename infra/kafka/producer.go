package kafka

import (
	"log"
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func NewKafkaConsumer(msgChan chan *ckafka.Message) *kafkaConsumer {
	return &kafkaConsumer{
		MsgChan: msgChan,
	}

}

func NewKafkaProducer() *ckafka.Producer {
	configmap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KafkaBootstrapServer"),
	}

	p, err := ckafka.NewProducer(configmap)
	if err != nil {
		log.Println(err.Error())
	}

	return p
}

func Publish(msg string, topic string, producer *ckafka.Producer) error {
	message := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: ckafka.PartitionAny},
		Value:          []byte(msg),
	}
	err := producer.Produce(message, nil)
	if err != nil {
		return err
	}
	return nil
}
