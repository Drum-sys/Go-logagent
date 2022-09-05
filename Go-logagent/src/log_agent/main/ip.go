package main

import (
	"fmt"
	"net"
)

var (
	localIpArry []string
)

func init()  {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(fmt.Sprintf("get local ip failed. %v", err))
		return
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				localIpArry = append(localIpArry, ipnet.IP.String())
			}
		}
	}
	fmt.Println(localIpArry)
}
