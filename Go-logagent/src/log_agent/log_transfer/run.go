package main

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
)

func Run() (err error) {
	partitionList, err := kafka_client.client.Partitions(kafka_client.topic)
	if err != nil {
		logs.Error("failed to get the list of partitions: %s", err)
		return
	}
	fmt.Println(partitionList)

	for partition := range partitionList {
		//加这个是因为怕主go程退出导致子go程退出，阻塞在wait()
		kafka_client.wg.Add(1)
		pc, errRet := kafka_client.client.ConsumePartition(kafka_client.topic, int32(partition), sarama.OffsetNewest)
		if errRet != nil {
			err = errRet
			logs.Error("failed to connect consumer partition %d: %s\n", partition, err)
			return
		}

		defer pc.AsyncClose()
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				logs.Debug("parition:%d, Offset:%d, Key:%s, Value:%s\n", msg.Partition, msg.Offset, msg.Key, msg.Value)
				err = sendToES(kafka_client.topic, string(msg.Value))
			}
			kafka_client.wg.Done()
		}(pc)
	}

	kafka_client.wg.Wait()
	kafka_client.client.Close()

	return

}

func sendToES(topic, data string) (err error) {
	ctx := context.Background()
	msg := &LogMessage{}
	msg.Topic = topic
	msg.Message = data


	ind, err := esClient.Index().
		Index(topic).
		BodyJson(msg).
		Do(ctx)
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(ind)
	return
}
