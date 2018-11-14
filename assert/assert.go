package main

import "fmt"

func assert(i interface{}){
	// thu ep ve kieu int
	// v la gia tri sau khi ep,
	// ok true la ep thanh cong
	// nguoc lai tra ve false
	v, ok := i.(int)
	fmt.Println(v, ok)
}

func findType(i interface{}){
	switch i.(type) {
	case string:
		fmt.Printf("I am a string and my value is %s\n", i.(string))
	case int:
		fmt.Printf("I am a int and my value is %d\n", i.(int))
	default:
		fmt.Printf("Unknows type\n")
	}
}

func main(){
	var s int = 56
	assert(s)
	var i string = "Steven Paul"
	assert(i)

	findType("Naveen")
	findType(77)
	findType(89.98)
}
