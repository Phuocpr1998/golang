package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	// neu thay doi giai tri cua mot phan tu thi
	// cac mang con lai cung bi anh huong theo
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}