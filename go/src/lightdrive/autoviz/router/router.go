package router

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"math"
	"math/rand"
	"net/http"
	"time"
)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
func Init() {

	// 入库前解析
	App.ServerMux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
			conn, err := upgrader.Upgrade(w, r, nil)
			if err != nil {
				log.Println(err)
				return
			}
			rr := rand.New(rand.NewSource(time.Now().UnixNano()))
			for {
				time.Sleep(time.Second/20)
				positions:=make([]float32,0)
				for i:=0;i<10;i++{
					p:=make([]float32,3)
					for j:=0;j<3;j++{
						p[j]=10*rr.Float32()*float32(math.Pow(-1,float64(i)))
					}
					positions=append(positions,p...)
				}
				conn.WriteJSON(positions)
				fmt.Println(positions)
			}

			//for {
			//	var msg Message                // Read in a new message as JSON and map it to a Message object
			//	err := conn.ReadJSON(&msg)
			//	if err != nil {
			//		log.Printf("error: %v", err)
			//		delete(clients, ws)
			//		break
			//	}
			//	// Send the newly received message to the broadcast channel
			//	broadcast <- msg        }

	})

	// 入库前解析
	App.ServerMux.HandleFunc("/e", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		rr := rand.New(rand.NewSource(time.Now().UnixNano()))
		for {
			time.Sleep(time.Second/20)
			positions:=make([][]float32,0)
			for i:=0;i<100;i++{
				p:=make([]float32,3)
				for j:=0;j<3;j++{
					p[j]=10*rr.Float32()*float32(math.Pow(-1,float64(i)))
				}
				positions=append(positions,p)
			}
			conn.WriteJSON(positions)
			fmt.Println(positions)
		}

		//for {
		//	var msg Message                // Read in a new message as JSON and map it to a Message object
		//	err := conn.ReadJSON(&msg)
		//	if err != nil {
		//		log.Printf("error: %v", err)
		//		delete(clients, ws)
		//		break
		//	}
		//	// Send the newly received message to the broadcast channel
		//	broadcast <- msg        }

	})
}
