package main


import (
	"fmt"
	"net/http"
	"github.com/gorilla/websocket"

)

// INITIALIZE
var upgrader = websocket.Upgrader{
	ReadBufferSize:1024,
	WriteBufferSize:1024,
}

// CREATE VARIABLE WEBSOCKET

var clients []websocket.Conn

func main(){
	// CREATE ENDPOIND FOR CONNECT WEBSOKCET
	http.HandleFunc("/echo",func(w http.ResponseWriter , r *http.Request){
		// INITIALIZE CONFIG
		conn,_ := upgrader.Upgrade(w,r,nil)

		clients = append(clients , *conn)

		// LOOP IF CLIENT SEND TO SERVER
		for{
			// READ MESSAGE FROM BROWSER
			msgType,msg,err := conn.ReadMessage()
			if err != nil{
				return
			}

			// PRINT MESSAGE IN YOU CONSOLE TERMINAL
			fmt.Printf("%s send: %s\n", conn.RemoteAddr(),string(msg))

			// LOOP IF MESSAGE FOUND AND SEND AGAIN TO CLIENT FOR
			// WRITE IN YOU BROWSER
			for _,client := range clients{
				if err = client.WriteMessage(msgType,msg);err !=nil{
					return 
				}
			}

		}

	})
	// SEND YOU HTML FILE FOR OPEN TO BROWSER
	http.HandleFunc("/",func(w http.ResponseWriter , r *http.Request){
		http.ServeFile(w,r,"index.html")
		// w,r IS WRITE AND DELETE YOU INDEX.HTML

	})
	println("You server run 8080")
	http.ListenAndServe(":8080",nil)

}