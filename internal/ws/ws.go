package ws

// package ws
//
// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"time"
//
// 	"github.com/gorilla/websocket"
// )
//
// const (
// 	// Time allowed to write a message to the peer.
// 	writeWait = 10 * time.Second
//
// 	// Time allowed to read the next pong message from the peer.
// 	pongWait = 60 * time.Second
//
// 	// Send pings to peer with this period. Must be less than pongWait.
// 	pingPeriod = (pongWait * 9) / 10
//
// 	// Maximum message size allowed from peer.
// 	maxMessageSize = 512
// )
//
// var (
// 	newline = []byte{'\n'}
// 	space   = []byte{' '}
// )
//
// var upgrader = websocket.Upgrader{
// 	ReadBufferSize:  1024,
// 	WriteBufferSize: 1024,
// 	CheckOrigin:     func(r *http.Request) bool { return true }, // Allow all connections
//
// }
//
// type Client struct {
// 	hub    *Hub
// 	conn   *websocket.Conn
// 	userID string // User's identifier
// } //
//
// func createClient(hub *Hub, w http.ResponseWriter, r *http.Request) {
// 	userID := r.URL.Query().Get("user_id")
// 	if userID == "" {
// 		http.Error(w, "user id is mandatory", http.StatusInternalServerError)
// 		return
// 	}
//
// 	fmt.Println("new websocket client created", userID)
//
// 	conn, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
//
// 	client := &Client{
// 		hub:    hub,
// 		conn:   conn,
// 		userID: userID,
// 	}
//
// 	client.hub.register <- client
//
// 	// go client.writePump()
// 	// go client.readPump()
//
// }
