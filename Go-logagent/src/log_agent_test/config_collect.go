package main

import (
	"fmt"
	"github.com/astaxie/beego/config"

)

func main() {
	conf, err := config.NewConfig("ini", "./logcollect.conf")
	if err != nil {
		fmt.Println("new config failed, err", err)
		return
	}

	port, err := conf.Int("server::port")
	if err != nil {
		fmt.Println("read server:port failed, err:", err)
		return
	}
	fmt.Println("Port:", port)


	ngnix := conf.String("http::ngnix")
	if err != nil {
		fmt.Println("read http::ngnix failed, err:", err)
		return
	}
	fmt.Println("Ngnix:", ngnix)



}
