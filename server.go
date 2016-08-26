
package main

import (
	"fmt"
	"gopkg.in/redis.v3"
)

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
		fmt.Println(message.Channel)
		fmt.Println(message.Payload)
	}
}

func main() {
	ConnectNewClient()
}
