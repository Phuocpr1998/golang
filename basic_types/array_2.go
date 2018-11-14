package main

import "fmt"

func main() {
	// khong can dinh truoc so luong
	a := []int{2, 3, 5, 7, 11, 13}

	// khai bao truoc so luong
	// b := [2]int{2, 3}

	fmt.Println(a)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct{
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}
