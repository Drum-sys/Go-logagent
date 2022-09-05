package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"strings"
	"sync"
)

var (
	wg sync.WaitGroup
)
func main() {
	consumer, err := sarama.NewConsumer(strings.Split("172.18.214.119:9092", ","), nil)
	if err != nil {
		fmt.Printf("failed to start consumer: %s", err)
		return
	}

	partitionList, err := consumer.Partitions("ngnix_log")
	if err != nil {
		fmt.Printf("failed to get the list of partitions: %s", err)
		return
	}
	fmt.Println(partitionList)

	for partition := range partitionList {
		wg.Add(1)
		pc, err := consumer.ConsumePartition("ngnix_log", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to connect consumer partition")
			return
		}

		defer pc.AsyncClose()
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("parition:%d, Offset:%d, Key:%s, Value:%s\n", msg.Partition, msg.Offset, msg.Key, msg.Value)
			}
			wg.Done()
		}(pc)
	}
	wg.Wait()
	consumer.Close()
}