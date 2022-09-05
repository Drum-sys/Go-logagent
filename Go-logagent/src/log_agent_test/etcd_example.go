package main

import (
	"fmt"
	etcdClient "go.etcd.io/etcd/client/v3"
	"time"

	"context"
)

func main() {
	cli, err := etcdClient.New(etcdClient.Config{
		Endpoints: []string{"localhost:2379" },
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}

	fmt.Println("connect success")
	defer cli.Close()

/*	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	_, err = cli.Put(ctx, "/log_agent/conf/", "sample_value")
	cancel()
	if err != nil {
		fmt.Println("put failed err", err)
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "/log_agent/conf/")
	cancel()

	if err != nil {
		fmt.Println("get failed err", err)
		return
	}
	for _, v := range resp.Kvs {
		fmt.Printf("%s : %s\n", v.Key, v.Value)
	}*/

	// 轮询监控节点，查看配置文件的变化，若有变化，将结果返回给与etcd建立连接的客户端
	for {
		rch := cli.Watch(context.Background(), "/log_agent/conf/")
		for wresp := range rch {
			for _, v := range wresp.Events {
				fmt.Printf("%s %q : %q\n", v.Type, v.Kv.Key, v.Kv.Value)
			}
		}

	}


}
