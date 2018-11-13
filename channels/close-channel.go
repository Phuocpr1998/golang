package main

import (
	"fmt"
	"time"
)

func main()  {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			// more will be false if jobs close
			j, more := <-jobs

			if more {
				fmt.Println("Received job", j)
			} else {
				fmt.Println("Received all joibs")
				done<- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++{
		jobs <- j
		fmt.Println("Sent job", j)
		<-time.After(time.Second)
	}

	close(jobs)
	fmt.Println("Sent all jobs")
	<-done
}
