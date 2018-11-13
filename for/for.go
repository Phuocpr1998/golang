package main

import "fmt"

func main(){
	sum := 0
	for i := 0; i< 10; i++ {
		sum += i	
	}
	
	// bien dem duoc dinh nghia ngoai vong for
	j := 0
	for j := 0; j < 10; j++ {
		sum += sum	
	}

	fmt.Println(sum)
}
