
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
	
	_, err := client.Subscribe("test1")
	if err != nil{
		fmt.Println(" error ")
	}
	fmt.Println("siii")

}

func main() {
	ConnectNewClient()
}
