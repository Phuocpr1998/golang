package main

import "fmt"

func add (x int, y int) int {
	return x + y;
}

func add2(x,y int) int {
	return x + y;
}

func main(){
	fmt.Println(add2(10,11));
}