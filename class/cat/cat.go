package cat

import "fmt"

type Cat struct {
	Name string
	Age int
	Weight float32
}


func (c Cat) ToString(){
	fmt.Printf("Name: %s Age %d Weight %g\n", c.Name, c.Age, c.Weight)
} 

func New(name string, age int, weight float32) Cat{
	c := Cat{name, age, weight}
	return c
}
