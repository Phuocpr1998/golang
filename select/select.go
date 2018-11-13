package main

import (
	"fmt"
	"time"
)

func main(){
	c1 := make(chan string)
	c2 := make(chan string)

	go func(){
		time.Sleep(1*time.Second)
		c1 <- "one"
	}()

	select {
	case msg1 := <-c1:
		fmt.Println("Received", msg1)
	case <-time.After(2 * time.Second):
		fmt.Println("timeout 1")
	}

	go func() {
		time.Sleep(3*time.Second)
		c2<-"two"
	}()

	select {
	case msg1 := <-c2:
		fmt.Println("Received", msg1)
	case <-time.After(2 * time.Second):
		fmt.Println("timeout 2")
	}
}

