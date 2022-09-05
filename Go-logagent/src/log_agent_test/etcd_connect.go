package main

import (
	"fmt"
	etcdClient "go.etcd.io/etcd/client/v3"
	"time"
)

func main() {
	cli, err := etcdClient.New(etcdClient.Config{
		Endpoints: []string{"localhost:2379", "localhost:22379", "localhost:32379" },
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}

	fmt.Println("connect success")
	defer cli.Close()



}
