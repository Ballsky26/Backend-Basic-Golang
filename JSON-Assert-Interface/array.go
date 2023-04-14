package main

import (
	"encoding/json"
	"fmt"
)

type User_array struct {
	FullName string `json:"Name"`
	Age      int
}

func main_array() {

	var jsonString = `[
{"Name": "john wick", "Age": 27}, {"Name": "ethan hunt", "Age": 32}
]`
	var data []User
	var err = json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("user 1:", data[0].FullName)
	fmt.Println("user 2:", data[1].FullName)
}
