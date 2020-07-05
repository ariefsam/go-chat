package httpapi

import (
	"bytes"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var ChannelHub map[string]*Hub

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	ChannelHub = map[string]*Hub{}
}

// upgrader.CheckOrigin = func(r *http.Request) bool { return true }” before ”conn, err := upgrader.Upgrade(w, r, nil)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func ListenChannelHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	log.Println("New Websocket Connection, ", vars["channelID"])
	channelID := vars["channelID"]
	hub := getChannelHub(channelID)

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
	client.hub.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.

	go client.writePump()

	defer func() {
		client.hub.unregister <- client
		client.conn.Close()
	}()
	client.conn.SetReadLimit(maxMessageSize)
	client.conn.SetReadDeadline(time.Now().Add(pongWait))
	client.conn.SetPongHandler(func(string) error { client.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := client.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		//client.hub.broadcast <- message
		log.Println("ada message", string(message))
	}

	return

}

func getChannelHub(channelID string) (hub *Hub) {
	if _, ok := ChannelHub[channelID]; !ok {
		hub = newHub()
		go hub.run()
		ChannelHub[channelID] = hub
	} else {
		hub = ChannelHub[channelID]
	}
	return
}

func reader(conn *websocket.Conn) {
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// print out that message for clarity
		log.Println(string(p))

	}
}
