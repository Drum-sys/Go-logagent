package main

import (
	//"context"
	//"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
)

type LogMessage struct {
	APP string `json:"app"`
	Topic string `json:"topic"`
	Message string `json:"message"`
}


var (
	esClient *elastic.Client
)

func InitES(addr string) (err error) {
	//ctx := context.Background()

	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(addr))

	if err != nil {
		fmt.Println("connect es failed, err:", err)
		return
	}
	esClient = client
	return
	/*fmt.Println("connect es success")



	data, _ := json.Marshal(tweet)
	js := string(data)

	res, err := client.Index().Index("twitter").Id("1").BodyJson(js).Do(ctx)
	if err != nil {
		panic(err)
		return
	}

	fmt.Println(res)
	fmt.Println("insert data into es success")*/
}
