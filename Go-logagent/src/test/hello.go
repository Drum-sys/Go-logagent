package main

import "time"

func main() {
	for i := 0; i<100; i++ {
		go TestPrint(i)
	}
	time.Sleep(time.Second * 1)
}