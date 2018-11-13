package main

import "fmt"

func split(sum int) (x, y int) {
	var s = sum	
	x = s*4/9
	y = s - x
	return
}

func main(){
	fmt.Println(split(17));
}
