package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

// truyen nhan data giua cac channel
func main() {
	var ops uint64
	for i:=0;i<50;i++{
		go func() {
			for {
				atomic.AddUint64(&ops, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	opsResult := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsResult)
}
