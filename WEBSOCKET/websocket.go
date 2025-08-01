package websocket

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader=websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}


func handlerequest(w http.ResponseWriter,r *http.Request){
	conn,err:=upgrader.Upgrade(w,r,nil)
	
	if err!=nil{
		log.Println("Connection failed: ",err.Error())
		return
	}
	defer conn.Close()

	log.Println("Connection established")

	for{
		mesagetype,message,err:=conn.ReadMessage()
		if err!=nil{
			log.Println("error reading message: ",err.Error())
			break
		}

		log.Println("Message received: ",string(message))

		msg:=fmt.Sprintf("hello %s from server",message)

		time.Sleep(2*time.Second)
		er:=conn.WriteMessage(mesagetype,[]byte(msg))
		if er!=nil{
			log .Println("error sending message",er.Error())
			break
		}
	}
	log.Println("connection lost")


}


func Controller(){
	r:=http.NewServeMux()

	r.HandleFunc("/websocket",handlerequest)

	fmt.Println("server started at port 8080")
	http.ListenAndServe(":8080",r)

}