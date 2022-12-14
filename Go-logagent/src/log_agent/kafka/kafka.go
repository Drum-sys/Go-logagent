package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
)

var (
	client sarama.SyncProducer
)
func InitKafka(addr string) (err error){
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal // 将日志发送给kafka，kafak发送ack确认收到消息
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 将日志信息分配到不同的机器上
	config.Producer.Return.Successes = true

	client, err = sarama.NewSyncProducer([]string{addr}, config)
	if err != nil {
		logs.Error("init kafka producer failed, err", err)
		return
	}

	logs.Debug("init kafka success")
	return
}

func SendKafka(data, topic string) (err error)  {
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)

	_, _, err = client.SendMessage(msg)
	if err != nil {
		logs.Error("send message failed, err:%v, data:%v, topic:%v", err, data, topic)
		return
	}
	//fmt.Println(pid, offset)
	//logs.Debug(" send success, pid：%v offset: %v, topic:%v\n", pid, offset, topic)
	return
}