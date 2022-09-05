package main

import (
	"errors"
	"fmt"
	"laonanhai/src/log_agent/tailf"

	//"github.com/Shopify/sarama"

	"github.com/astaxie/beego/config"
)

type Config struct {
	logLevel string
	logPath string
	chanSize int
	kafkaAddr string
	collectConf []tailf.CollectPath
	etcdAddr string
	etcdKey string
}

var (
	appConf *Config
)

func LoadCollectConf(conf config.Configer) (err error) {
	var cc tailf.CollectPath
	cc.LogPath = conf.String("collect::log_path")
	if len(cc.LogPath) == 0 {
		err = errors.New("invalid collect::log_path")
		return
	}

	cc.Topic = conf.String("collect::topic")
	if len(cc.Topic) == 0 {
		err = errors.New("invalid collect::topic")
		return
	}
	appConf.collectConf = append(appConf.collectConf, cc)

	return
}

func LoadConf(confType, filename string) (err error) {
	conf, err := config.NewConfig(confType, filename)
	if err != nil {
		fmt.Println("read conf failed, err:", err)
		return
	}

	appConf = &Config{}

	appConf.logLevel = conf.String("logs::log_level")
	if len(appConf.logPath) == 0 {
		appConf.logLevel = "debug"
	}

	appConf.logPath = conf.String("logs::log_path")
	if len(appConf.logPath) == 0 {
		appConf.logPath = "./logs"
	}

	appConf.chanSize, err = conf.Int("collect::chan_size")
	if err != nil {
		appConf.chanSize = 100
	}

	appConf.kafkaAddr = conf.String("kafka::server_addr")
	if len(appConf.kafkaAddr) == 0 {
		err = fmt.Errorf("invalid kafka addr")
		return
	}

	appConf.etcdAddr = conf.String("etcd::addr")
	if len(appConf.etcdAddr) == 0 {
		err = fmt.Errorf("invalid etcd addr")
		return
	}

	appConf.etcdKey = conf.String("etcd::configKey")
	if len(appConf.etcdKey) == 0 {
		err = fmt.Errorf("invalid etcd key")
	}


	err = LoadCollectConf(conf)
	if err != nil {
		fmt.Println("load collect conf, err:", err)
		return
	}
	return
}
