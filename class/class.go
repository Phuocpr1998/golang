package main

//import "./cat"
import (
	"./rectangle" 
	"fmt"
	)

func main(){
	/// Class Cat
	//c := cat.New("Kitty", 10, 30)
	//c.ToString()
	
	/// class Rectangle
	r := rectangle.Rectangle{
		Name: "Sharp 1",		
		Width: 10,
		Height: 5,
	}

	r.ToString();	

	fmt.Printf("Dien tich: %g\n", r.Area())
}
