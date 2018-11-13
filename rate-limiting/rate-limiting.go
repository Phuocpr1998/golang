package main

import (
	"fmt"
	"time"
)

func main(){
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++{
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200*time.Millisecond)

	for req := range requests{
		<-limiter
		fmt.Println("Request", req, time.Now())
	}

	burstryLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstryLimiter <- time.Now()
	}

	go func() {
		for t := range  time.Tick(200*time.Millisecond){
			burstryLimiter<-t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1;i <=5; i++{
		burstyRequests <- i
	}

	close(burstyRequests)

	for req := range burstyRequests{
		<-burstryLimiter
		fmt.Println("Request", req, time.Now())
	}

}
