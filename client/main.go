package main

import (
	"context"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, _, err := websocket.DefaultDialer.DialContext(ctx, "ws://localhost:8080/ws", nil)
	if err != nil {
		log.Fatal("Failed to connect:", err)
	}
	defer conn.Close()

	log.Println("Connected successfully")

	// Do nothing, just stay connected
	time.Sleep(10 * time.Second)

	log.Println("Disconnecting")
}
