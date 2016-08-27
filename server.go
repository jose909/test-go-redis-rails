
package main

import (
	"fmt"
	"gopkg.in/redis.v3"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"net/http"
)

type Request struct{
	Id int
	Name string
}

type Client struct{
	Id int
	websocket *websocket.Conn
}

var Clients = make(map[int]Client)
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
	go ConnectNewClient()

	mux := mux.NewRouter()
	mux.HandleFunc("/Subscribe/", Subscribe).Methods("GET")
	http.Handle("/", mux)
	fmt.Println("El servidor se encuentra en el puerto 8000")
	http.ListenAndServe(":8000", nil)
}
func Subscribe(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Upgrade(w,r,nil,1024, 1024)
	if err !=nil{
		return
	}
	fmt.Println("Nuevo web socket")
	count := len(Clients)
	new_client := Client{count,ws}
	Clients[count] = new_client
	fmt.Println("Se fue client")

	for{
		_, _, err := ws.ReadMessage()
		if err != nil{
			delete(Clients, new_client.Id)
			fmt.Println("Se fue client")
			return
		}
	}
}