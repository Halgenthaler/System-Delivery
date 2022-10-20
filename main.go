package main

import (
	"fmt"
	"log"

	"github.com/Halgenthaler/System-Delivery/infra/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error .env file")
	}
}

func main() {
	// TEST TO READ THE MESSAGE
	//producer := kafka.NewKafkaProducer()
	//kafka.Publish("Hello World", "readtest", producer)

	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume()

	for msg := range msgChan {
		fmt.Println(string(msg.Value))
	}

	//---------------------------------------------------------
	// FUNC TO TEST THE SIMULATOR
	//	route := route.Route{
	//		ID:       "1",
	//		ClientID: "1",
	//	}
	//	route.LoadPositions()
	//	stringjson, _ := route.ExportJsonPositions()
	//	fmt.Println(stringjson[2:8])
	//---------------------------------------------------------
}
