
package main

import (
	"fmt"
	"gopkg.in/redis.v3"
	"encoding/json"
)

type Request struct{
	Id int
	Name string
}

func ConnectNewClient() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
		})
	
	pubsub, err := client.Subscribe("test1")
	if err != nil{
		fmt.Println(" error ")
	}
	for{
		message, err := pubsub.ReceiveMessage()
		if err != nil{
			fmt.Println("No es posible leer el mensaje")
		}

		request := Request{}
		if err := json.Unmarshal([]byte(message.Payload), &request); err !=nil{
			fmt.Println("No es posible leer el json")
		}
		fmt.Println(request.Name)
		fmt.Println(request.Id)
		//fmt.Println(message.Channel)
		//fmt.Println(message.Payload)
	}
}

func main() {
	ConnectNewClient()
}
