package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
	client, err := sarama.NewClient([]string{"localhost:9092"}, nil)
	if err != nil {
		panic(err)
	}
	topics, err := client.Topics()
	if err != nil {
		panic(err)
	}
	fmt.Println(topics)
	sarama.NewConsumerGroup()
	group, err := sarama.NewConsumerGroupFromClient("1", client)
	if err != nil {
		panic(err)
	}

}
