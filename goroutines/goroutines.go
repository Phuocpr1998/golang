/*package main

import "fmt"

func f(from string){
	for i:=0; i<3; i++{
		fmt.Println(from, ":", i)	
	}
}

func main(){
	f("direct")
 	go f("goroutine")
	
	go func(msg string) {
		fmt.Println(msg)
	}("going")
	
	fmt.Scanln()
	fmt.Println("done")
}*/	

package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}
