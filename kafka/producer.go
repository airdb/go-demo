package main

import (
	"fmt"
	_ "os"
	_ "os/signal"
	"strconv"

	"github.com/Shopify/sarama"
	_ "github.com/bsm/sarama-cluster" //support automatic consumer-group rebalancing and offset tracking
)

var (
	brokerList = []string{"broker1:9092", "broker2:9092", "broker3:9092"}
	topic      = "RMB"
	maxRetry   = 5
)

func producer() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = maxRetry
	config.ClientID = "go-kafka-consumer"
	config.Consumer.Return.Errors = true
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			panic(err)
		}
	}()
	for i := 1; i < 100; i++ {
		msg := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder("dean Cool " + strconv.Itoa(i)),
		}
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)
	}
}

func consumer() {
	messageCountStart := 0

	config := sarama.NewConfig()
	config.ClientID = "go-kafka-consumer"
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.CommitInterval = 1
	// config.Consumer.Offsets.Reset = true
	// config.Consumer.Offsets.Initial = sarama.OffsetNewest
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	client, err := sarama.NewClient(brokerList, config)
	fmt.Println(client.GetOffset(topic, 0, sarama.OffsetOldest))

	brokers := brokerList
	master, err := sarama.NewConsumer(brokers, config)

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := master.Close(); err != nil {
			panic(err)
		}
	}()

	partitionList, err := master.Partitions(topic)
	if err != nil {
		fmt.Println("Failed to get the list of partitions: ", err)
	}

	for partition := range partitionList {
		fmt.Println("partitionList====", partition)
		// pc, err := consumer.ConsumePartition("topic.ops.falcon", int32(partition), sarama.OffsetNewest)

		// consumer, err := master.ConsumePartition(topic, 7, sarama.OffsetOldest)
		consumer, err := master.ConsumePartition(topic, int32(partition), sarama.OffsetOldest)
		// consumer, err := master.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		// consumer, err := master.ConsumePartition(topic, 0, sarama.OffsetNewest)
		if err != nil {
			panic(err)
		}
		for i := 0; i < 10; i++ {
			messageCountStart++
			msg := <-consumer.Messages()
			// _ consumer.MarkOffset(msg, "")
			// fmt.Println(string(msg.Key), string(msg.Value), partition, msg.Offset)
			fmt.Printf("offset: %d, partition: %d, msg: %s\n", msg.Offset, partition, string(msg.Value))
			// fmt.Println(msg.Timestamp)
			// consumer.CommitOffset(msg.Offset)
		}
		// signals := make(chan os.Signal, 1)
		// signal.Notify(signals, os.Interrupt)
		// doneCh := make(chan struct{})
		// go func() {
		// 	for {
		// 		select {
		// 		case err := <-consumer.Errors():
		// 			fmt.Println(err)
		// 		case msg := <-consumer.Messages():
		// 			messageCountStart++
		// 			fmt.Println("Received messages", string(msg.Key), string(msg.Value), partition)
		// 		case <-signals:
		// 			fmt.Println("Interrupt is detected")
		// 			doneCh <- struct{}{}
		// 		}
		// 	}
		// }()
		// <-doneCh
		consumer.Close()
		fmt.Println("Processed", messageCountStart, "messages")
		break
		messageCountStart = 0
	}

}

func main() {
	// consumer()
	producer()
}
