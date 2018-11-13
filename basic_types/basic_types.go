package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe bool = false
	MaxInt uint64 = 1<<64 -1
	z	complex128 = cmplx.Sqrt(-5 + 12i)
	ZeroValue int // tu dong gan bang 0 neu khong khoi tao truoc
		      // Bang "" doi voi string va bang false voi type bool
)

func main(){

	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[:4] // lay cac phan tu trong mang tu 0 -> 3
	fmt.Println(s)

	s = primes[4:] // lay cac phan tu trong mang tu 4 -> n
	fmt.Println(s)

	s = primes[:] // lay tat ca cac phan tu cua mang
	fmt.Println(s)
}
