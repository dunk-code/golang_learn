package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"sync"
)

func main() {
	topic := "test"

	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()
	partitions, err := consumer.Partitions(topic)
	if err != nil {
		panic(err)
	}
	fmt.Println(partitions)
	wg := sync.WaitGroup{}
	for partition := range partitions {
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			panic(err)
		}
		defer pc.AsyncClose()
		wg.Add(1)
		go func(partitionConsumer sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%v Value:%s \n", msg.Partition, msg.Offset, msg.Key, string(msg.Value))
			}
			wg.Done()
		}(pc)
	}
	wg.Wait()
}
