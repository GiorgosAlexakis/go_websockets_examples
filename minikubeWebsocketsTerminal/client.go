package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gorilla/websocket"
)

func main() {
	ca, err := ioutil.ReadFile("/home/george/.minikube/ca.crt")
	if err != nil {
		log.Fatal(err)
	}
	// read client cert
	clientCert, err := ioutil.ReadFile("/home/george/.minikube/profiles/minikube/client.crt")
	if err != nil {
		log.Fatal("Error loading client cert", err)
	}
	// read client key
	clientKey, err := ioutil.ReadFile("/home/george/.minikube/profiles/minikube/client.key")
	if err != nil {
		log.Fatal("Error loading client key", err)
	}
	value1 := "ca: " + string(ca)
	value2 := "cert: " + string(clientCert)
	value3 := "key: " + string(clientKey)

	dialer := websocket.DefaultDialer // use default dialer
	dialer.Subprotocols = []string{value1, value2, value3}
	//dialer.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	url := "ws://192.168.49.2:30110/api/v1/namespaces/default/pods/hello-minikube-6ddfcc9757-g4484/exec?command=echo&command=ls&stderr=true&stdout=true"
	//dial websocket

	c, _, err := dialer.Dial(url, nil)
	fmt.Println(c)
	if err != nil {
		log.Fatal("dial:", err)
	}
	// receive websocket message
	defer c.Close()
	for {
		_, message, err := c.NextReader()
		//_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		// print message
		//log.Printf("recv: %s", message)
		fmt.Println(message)
	}

}
