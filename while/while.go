package main

import "fmt"

func main(){
	sum := 1
	// vong lap while
	for sum < 1000 {
		sum += sum	
	}

	/// vong lap vo han
	for{
		sum += sum
		break	// ngat vong lap
	}
	
	fmt.Println(sum) // 1024
}
