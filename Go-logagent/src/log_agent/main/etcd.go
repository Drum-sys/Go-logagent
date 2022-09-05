package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"go.etcd.io/etcd/api/v3/mvccpb"
	etcdClient "go.etcd.io/etcd/client/v3"
	"laonanhai/src/log_agent/tailf"
	"strings"
	"time"
)

type EtcdClient struct {
	client *etcdClient.Client
	keys []string
}
var (
	etcd_Client *EtcdClient
)
func InitEtcd(addr string, key string) (collectConf []tailf.CollectPath, err error) {
	cli, err := etcdClient.New(etcdClient.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		logs.Error("connect etcd failed, err:", err)
		return
	}
	etcd_Client = &EtcdClient{
		client: cli,
	}
	if strings.HasSuffix(key, "/") == false {
		key = key + "/"
	}

	for _, localIp := range localIpArry {
		etcdKey := fmt.Sprintf("%s%s", key, localIp)
		etcd_Client.keys = append(etcd_Client.keys, etcdKey)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		resp, err := cli.Get(ctx, etcdKey)
		if err != nil {
			logs.Error("client get from etcd failed, err:%v", err)
			continue
		}
		for _, v := range resp.Kvs {
			if string(v.Key) == etcdKey {
				err = json.Unmarshal(v.Value, &collectConf)
				if err != nil {
					logs.Error("unmarshall failed, err:", err)
					continue
				}
				logs.Debug("get conf from etcd success, log config is:", collectConf)
			}
		}

		cancel()
	}

	InitWtcdWatcher()
	return
}

func InitWtcdWatcher() {
	for _, key := range etcd_Client.keys{
		go WatchKey(key)
	}
}

func WatchKey(key string) {
	cli, err := etcdClient.New(etcdClient.Config{
		Endpoints: []string{"localhost:2379" },
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}
	logs.Debug("start watch key:", key)
	for {
		rch := cli.Watch(context.Background(),  key)
		var collectConf []tailf.CollectPath
		var getConfSucc bool
		getConfSucc = true
		for wresp := range rch {
			for _, v := range wresp.Events {
				if v.Type == mvccpb.DELETE {
					logs.Warn("key[%s] config delete", key)
					continue
				}
				if v.Type == mvccpb.PUT && string(v.Kv.Key) == key{
					err = json.Unmarshal(v.Kv.Value, &collectConf)
					if err != nil {
						logs.Error("key [%s], Unmarshall [%s], err:%v", err)
						getConfSucc = false
						continue
					}
				}
				logs.Debug("get config from etcd %s %q : %q\n", v.Type, v.Kv.Key, v.Kv.Value)
			}
			if getConfSucc {
				logs.Debug("get config from etcd success, %v", collectConf)
				tailf.UpdateConf(collectConf)
			}
		}

	}
}