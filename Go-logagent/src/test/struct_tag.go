package main

import (
	"encoding/json"
	"fmt"
)

type Student1 struct {
	Name  string `json:"Student_name"`
	Age int `json:"age"`
	Score int `json:"score"`
}

func main() {
	var stu Student1 = Student1{
		Name: "ljw",
		Age: 15,
		Score: 98,
	}

	data, err := json.Marshal(stu)
	if err != nil {
		fmt.Println("json encode failed")
		return
	}

	fmt.Println(string(data))

}
