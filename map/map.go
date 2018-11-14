package main

import "fmt"

func main(){
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("Map", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	delete(m, "k2")

	// kiem tra xem con phan tu co key "k2" ko
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{
		"foo" : 1,
		"bar" : 2,
	}
	fmt.Println("Map:", n)

}
