package main

import "fmt"

func main(){
	// create channel with buffer
	messages := make(chan string, 2)
	
	go func(){messages <- "world"}()
	go func(){messages <- "hello"}()
	
	// messages <- "hello"	
	// messages <- "world"

	fmt.Println(<-messages)
    	fmt.Println(<-messages)
}
