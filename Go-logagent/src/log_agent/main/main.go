package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"laonanhai/src/log_agent/kafka"
	"time"

	"laonanhai/src/log_agent/tailf"
)

func main() {

	filename := "./conf/logagent.conf"
	confType := "ini"

	// 初始化配置文件
	err := LoadConf(confType, filename)
	if err != nil {
		fmt.Println("load conf err:", err)
		panic("load conf failed")
		return
	}
	logs.Debug("load conf success:%v", appConf)
	// 初始化日志
	err = InitLogger()
	if err != nil {
		fmt.Println("load logger err:", err)
		panic("load logger failed")
		return
	}
	logs.Debug("init logger success")

	collectConf, err := InitEtcd(appConf.etcdAddr, appConf.etcdKey)
	if err != nil {
		logs.Error("init etcd failed, err:%v", err)
		return
	}
	logs.Debug("init etcd success")

	err = tailf.InitTail(collectConf, appConf.chanSize)
	if err != nil {
		logs.Error("init tail failed, err:%v", err)
		return
	}
	logs.Debug("init tail success:%v", appConf)

	err = kafka.InitKafka(appConf.kafkaAddr)
	if err != nil {
		logs.Error("init kafka failed, err:%v", err)
		return
	}

	logs.Debug("initialize all success")

	go func() {
		var count int
		for {
			count++
			logs.Debug("test for logger %d", count)
			time.Sleep(time.Millisecond * 1000)
		}
	}()
	err = ServerRun()
	logs.Debug("end=====")
	if err != nil {
		logs.Error("serverRun failed, err:%v", err)
		return
	}

	logs.Info("program exited")
}
