
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


func ConnectNewClient(request_chanel chan Request) {
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

		request_chanel <- request
		//fmt.Println(message.Channel)
		//fmt.Println(message.Payload)
	}
}

func main() {
	Channel_request := make (chan Request)
	go ConnectNewClient(Channel_request)
	go ValidateChanel(Channel_request)

	mux := mux.NewRouter()
	mux.HandleFunc("/subscribe/", Subscribe).Methods("GET")
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

func ValidateChanel(request chan Request){
	for{
		select{
			case r := <- request:
			SendMessage(r)//Enviar mensaje
		}
	}
}

func SendMessage(request Request) {
	for _, client := range Clients{
		if err := client.websocket.WriteJSON(request); err!= nil{
			return
		}
	}
	
}