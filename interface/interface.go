package main

import "fmt"

type Shap  interface{
	Area() float64
}


type Rect struct {
	height int
	width int
}

type Square struct {
	width int
}

func (r Rect) Area() float64{
	return float64(r.height*r.width)
}

func (s Square) Area() float64{
	return float64(s.width*s.width)
}

func main(){
	rect := Rect{
		height:10,
		width:5,
	}

	square := Square{
		5,
	}

	sharps := []Shap{rect, square}
	for i, v := range sharps{
		fmt.Println("Are ", i , ":", v.Area())
	}

}

