package main

import (
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
	"strings"
	"sync"

	//"sync"
)

/*var (
	wg sync.WaitGroup
	//kafkaChan chan string
)*/
type kafkaClient struct {
	client sarama.Consumer
	addr string
	topic string
	wg sync.WaitGroup
}

var (
	kafka_client *kafkaClient
)
func InitKafka(addr, topic string) (err error) {

	kafka_client = &kafkaClient{}

	consumer, err := sarama.NewConsumer(strings.Split(addr, ","), nil)
	if err != nil {
		logs.Error("init kafka failed, err:%v", err)
		return
	}
	kafka_client.client = consumer
	kafka_client.addr = addr
	kafka_client.topic = topic
	return

	/*partitionList, err := consumer.Partitions(topic)
	if err != nil {
		logs.Error("failed to get the list of partitions: %s", err)
		return
	}
	fmt.Println(partitionList)

	for partition := range partitionList {
		//wg.Add(1)
		pc, errRet := consumer.ConsumePartition("ngnix_log", int32(partition), sarama.OffsetNewest)
		if errRet != nil {
			err = errRet
			logs.Error("failed to connect consumer partition %d: %s\n", partition, err)
			return
		}

		defer pc.AsyncClose()
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				logs.Debug("parition:%d, Offset:%d, Key:%s, Value:%s\n", msg.Partition, msg.Offset, msg.Key, msg.Value)
			}
			//wg.Done()
		}(pc)
	}
	//wg.Wait()
	consumer.Close()*/
}