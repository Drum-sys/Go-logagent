package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"laonanhai/src/log_agent/kafka"
	"laonanhai/src/log_agent/tailf"
	"time"
)

func ServerRun() (err error) {
	for {
		msg := tailf.GetOneLine()
		fmt.Println(msg)
		err = SendMsgKafka(msg)
		if err != nil {
			logs.Error("send to kafka failed, err:%v", err)
			time.Sleep(time.Second)
			continue
		}
	}
	return
}

func SendMsgKafka(msg *tailf.TextMsg) (err error) {
	//logs.Debug("read msg:%s, topic:%s", msg.Msg, msg.Topic)
	kafka.SendKafka(msg.Msg, msg.Topic)
	fmt.Println(msg.Msg, msg.Topic)
	return

}
