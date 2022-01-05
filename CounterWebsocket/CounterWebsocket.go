package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var mu sync.Mutex
var count int

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func counter(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ws upgrader: %v\n", err)
		return
	}
	mu.Lock()

	// write to the client the count value every second
	for {
		mu.Unlock()
		err = conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%d", count)))
		if err != nil {
			fmt.Fprintf(os.Stderr, "write to client: %v\n", err)
		}
		mu.Lock()
		// sleep for 1 second
		time.Sleep(1 * time.Second)
		count++
	}

}
func main() {
	http.HandleFunc("/counter", counter)
	log.Fatal(http.ListenAndServe("localhost:12345", nil))
}
