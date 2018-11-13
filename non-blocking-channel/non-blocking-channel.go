package main

import (
	"fmt"
	"time"
)

func main(){
	messages := make(chan string)
	signals := make(chan bool)
	go func() {messages<-"Hi"}()

	select {
	case msg:=<-messages:
		fmt.Println("Received message", msg)
	case <-time.After(2 * time.Second):
		fmt.Println("No message received")
	}

	msg:="hi"

	select {
	case messages<-msg:
		fmt.Println("Sent message", msg)
	default:
		fmt.Println("No message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("Received message", msg)
	case sig := <-signals:
		fmt.Println("Received signal", sig)
	default:
		fmt.Println("No activity")
	}

}