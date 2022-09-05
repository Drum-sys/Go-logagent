package main

import (
	"fmt"
	"github.com/astaxie/beego/config"
)

type LogConf struct {
	kafkaAddr string
	ESAddr string
	LogPath string
	LogLevel string
	kafkaTopic string
}

var (
	logConf *LogConf
)

func InitConf(confType, filename string) (err error) {
	conf, err := config.NewConfig(confType, filename)
	if err != nil {
		fmt.Println("read conf failed, err:", err)
		return
	}

	logConf = &LogConf{}

	logConf.LogLevel = conf.String("logs::log_level")
	if len(logConf.LogLevel) == 0 {
		logConf.LogLevel = "debug"
	}

	logConf.LogPath = conf.String("logs::log_path")
	if len(logConf.LogPath) == 0 {
		logConf.LogPath = "./logs"
	}

	logConf.kafkaAddr = conf.String("kafka::server_addr")
	if len(logConf.kafkaAddr) == 0 {
		err = fmt.Errorf("invalid kafka addr")
		return
	}

	logConf.kafkaTopic = conf.String("kafka::topic")
	if len(logConf.kafkaTopic) == 0 {
		err = fmt.Errorf("invalid kafka topic")
		return
	}

	logConf.ESAddr = conf.String("es::addr")
	if len(logConf.ESAddr) == 0 {
		err = fmt.Errorf("invalid es addr")
		return
	}
	return
}