package main

import "fmt"

func main(){
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	// close channel
	// neu khong o vong lap duoi se lap vo han
	close(queue)

	//lay tung phan tu chua trong channel queue
	for elem := range queue{
		fmt.Println(elem)
	}

	// array
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}

	// _ bo qua gia tri index
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
