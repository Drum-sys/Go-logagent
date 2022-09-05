package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/jmoiron/sqlx"
	clientv3 "go.etcd.io/etcd/client/v3"
	model "logadmin/models"
	_ "logadmin/routers"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego"
)

func initDb() (err error) {
	database, err := sqlx.Open("mysql", "root:19981015@tcp(127.0.0.1:3306)/test")
	if err != nil {
		logs.Warn("open mysql failed,", err)
		return
	}

	model.InitDb(database)
	return
}

func initEtcd() (err error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}

	model.InitEtcd(cli)
	return
}

func main() {
	err := initDb()
	if err != nil {
		logs.Warn("initDb failed, err:%v", err)
		return
	}

	err = initEtcd()
	if err != nil {
		logs.Warn("init etcd failed, err:%v", err)
		return
	}
	beego.Run()
}
