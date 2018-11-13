package rectangle


import "fmt"

type Rectangle struct {
	Name string
	Width, Height int
}

func (r Rectangle) ToString(){
	fmt.Printf("Name %s Width %d, Height %d\n", r.Name, r.Width, r.Height)
}

func (r Rectangle) Area() float32{
	return float32(r.Width * r.Height);
}


