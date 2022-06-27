package main

import (
	"github.com/Shopify/sarama"
	"log"
	"time"
)

type Broker struct {
	Addrs   []string
	Version string
}

type KafkaConf struct {
	Broker   Broker
	Producer map[string]string
}

type MsgProducer struct {
	producer sarama.AsyncProducer
	topic    string
}

func (p MsgProducer) asyncLoopResp() {
	log.Printf("MsgProducer start AsyncLoopResp")
	for {
		select {
		case <-p.producer.Successes():
			log.Println("AsyncLoopResp get changelog success")
		case err := <-p.producer.Errors():
			log.Println("AsyncLoopResp get changelog rsp error", err)
		}
	}
}

func main() {
	kafkaConf := KafkaConf{
		Broker: Broker{
			Addrs:   []string{"localhost:9092"},
			Version: "1",
		},
		Producer: map[string]string{
			"test": "test",
		},
	}
	p := newMsgProducer(kafkaConf)
	go p.asyncLoopResp()
	p.producer.Input() <- &sarama.ProducerMessage{
		Topic: p.topic,
		Key:   sarama.StringEncoder("key"),
		Value: sarama.StringEncoder("this is a msg"),
	}
	time.Sleep(time.Second * 5)
}

func newMsgProducer(conf KafkaConf) *MsgProducer {
	broker := conf.Broker
	topic := conf.Producer["test"]
	p := newAsyncProducer(broker)
	producer := &MsgProducer{
		producer: p,
		topic:    topic,
	}
	return producer
}

func newAsyncProducer(broker Broker) sarama.AsyncProducer {
	// version, err := sarama.ParseKafkaVersion(broker.Version)
	config := sarama.NewConfig()
	// config.Version = version

	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回
	config.Producer.Return.Errors = true

	client, err := sarama.NewAsyncProducer(broker.Addrs, config)
	if err != nil {
		panic(err)
	}
	return client
}
