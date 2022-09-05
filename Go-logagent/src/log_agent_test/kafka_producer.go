package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

//从文件读日志
func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal // 将日志发送给kafka，kafak发送ack确认收到消息
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 将日志信息分配到不同的机器上
	config.Producer.Return.Successes = true

	client, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		fmt.Println("producer close, err", err)
		return
	}

	defer client.Close()
	for {

		msg := &sarama.ProducerMessage{}
		msg.Topic = "ngnix_log"
		msg.Value = sarama.StringEncoder("this is a good test, my message is good")

		pid, offset, err := client.SendMessage(msg)
		if err != nil {
			fmt.Println("send message failed, err", err)
			return
		}
		fmt.Printf("pid：%v offset: %v\n", pid, offset)
		time.Sleep(10*time.Second)
	}


}
