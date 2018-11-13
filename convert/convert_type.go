package main

import (
	"fmt"
	"math"
)

func main(){
	var x, y int = 3,4
	var f float64 = math.Sqrt(float64(x*x + y*y)) // ep kieu int ve kieu float64
	var z uint = uint(f)
	fmt.Println(x, y, z)	
}
