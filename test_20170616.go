package go_discovery

import "fmt"

package main

import "fmt"

func fac(n int) int {
	if n <= 0 {
		return 1
	}
	return n * fac(n-1)
}

func facIter(n int) int {
	result := 1
	for n > 0 {
		result *= n
		n--
	}
	return result
}

func fac2(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		fmt.Println("i is : ", i)
		result *= i
	}
	return result
}



func main() {
	fmt.Println("This is fac ", fac(4))
	fmt.Println("This is facIter ", facIter(4))
	fmt.Println("This is fac2 ", fac2(4))
	//	var a = 10
	//	b := "little"
	//	b += " gophers"
	//	fmt.Println(a, b)
}
