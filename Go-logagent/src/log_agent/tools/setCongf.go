package main

import (
	"encoding/json"
	"fmt"
	etcdClient "go.etcd.io/etcd/client/v3"
	"laonanhai/src/log_agent/tailf"
	"time"

	"context"
)

const (
	EtcdKey = "/oldboy/backend/logagent/config/172.18.214.119"
)

func SetLogConfToEtcd() {
	cli, err := etcdClient.New(etcdClient.Config{
		Endpoints: []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}

	fmt.Println("connect success")
	defer cli.Close()

	var logConfArr []tailf.CollectPath
	logConfArr = append(logConfArr, tailf.CollectPath{
		LogPath: "/home/drum/go_workspace/src/laonanhai/src/log_agent/logs/logagent.log",
		Topic: "ngnix_log",
	})

	logConfArr = append(logConfArr, tailf.CollectPath{
		LogPath: "/home/drum/go_workspace/src/laonanhai/src/log_agent/logs/error.log",
		Topic: "ngnix_log_err",
	})
	data, err := json.Marshal(logConfArr)
	if err != nil {
		fmt.Println("json etcd data failed")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	//cli.Delete(ctx, EtcdKey)
	//return
	_, err = cli.Put(ctx, EtcdKey, string(data))
	cancel()
	if err != nil {
		fmt.Println("put failed err", err)
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, EtcdKey)
	cancel()
	if err != nil {
		fmt.Println("get failed err", err)
		return
	}
	for _, v := range resp.Kvs {
		fmt.Printf("%s : %s\n", v.Key, v.Value)
	}
}

func main() {
	SetLogConfToEtcd()
}