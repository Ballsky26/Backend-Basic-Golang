package main

import (
	"encoding/json"
	"fmt"
)

type User_json struct {
	FullName string `json:"Name"`
	Age      int
}

func main_json() {
	var jsonString = `{"Name":"Ballsky","Age":27}`
	var jsonData = []byte(jsonString)
	var data User
	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("user :", data.FullName)
	fmt.Println("age :", data.Age)
}
