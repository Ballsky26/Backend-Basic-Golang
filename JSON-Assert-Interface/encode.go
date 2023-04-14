package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	FullName string `json:"Name"`
	Age      int
}

func main_() {
	var object = []User{{"Ballsky", 23}, {"Ari", 25}}
	var jsonData, err = json.Marshal(object)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var jsonString = string(jsonData)
	fmt.Println(jsonString)
}
