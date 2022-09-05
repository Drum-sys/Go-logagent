package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
)

type Tweet struct {
	User	string `json:"user"`
	Message string `json:"message"`
}

func main() {
	ctx := context.Background()

	client,err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://127.0.0.1:9200"))

	if err != nil {
		fmt.Println("connect es failed, err:", err)
		return
	}

	fmt.Println("connect es success")
	tweet := Tweet{
		User: "ljw",
		Message: "Take five",

	}
	data, _ := json.Marshal(tweet)
	js := string(data)

	res, err := client.Index().Index("twitter").Id("1").BodyJson(js).Do(ctx)
	if err != nil {
		panic(err)
		return
	}

	fmt.Println(res)
	fmt.Println("insert data into es success")
}
