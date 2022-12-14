package main

import (
	"fmt"
	"github.com/olivere/elastic/v7"

	"context"
)

type Student struct {
	Name         string  `json:"name"`
	Age          int64   `json:"age"`
	AverageScore float64 `json:"average_score"`
}

func GetESClient() (*elastic.Client, error) {

	client, err :=  elastic.NewClient(elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))

	fmt.Println("ES initialized...")

	return client, err

}

func main() {

	ctx := context.Background()
	esclient, err := GetESClient()
	if err != nil {
		fmt.Println("Error initializing : ", err)
		panic("Client fail ")
	}

	//creating student object
	newStudent := Student{
		Name:         "Gopher doe",
		Age:          10,
		AverageScore: 99.9,
	}

	//dataJSON, err := json.Marshal(newStudent)
	//js := string(dataJSON)
	ind, err := esclient.Index().
		Index("students").
		BodyJson(newStudent).
		Do(ctx)

	if err != nil {
		panic(err)
	}
	fmt.Println(ind)
	fmt.Println("[Elastic][InsertProduct]Insertion Successful")

}
