package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gorilla/websocket"
)

var ErrNoPath = errors.New("path required")

func execInput(input string, ws *websocket.Conn) {
	// Remove the newline character.
	input = strings.TrimSuffix(input, "\n")
	// write input to websocket
	err := ws.WriteMessage(websocket.TextMessage, []byte(input))
	if err != nil {
		log.Println("write:", err)
	}
	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		fmt.Printf("%s\n", message)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	// read first argument
	url := os.Args[1]
	log.Printf("connecting to %s", url)
	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	//When the program closes close the connection
	defer c.Close()
	for {
		fmt.Print("$ ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if input == "\n" {
			fmt.Println("You must enter a command to get an output")
			continue
		}
		execInput(input, c)
		// print each line of cmdOutput

	}
}
