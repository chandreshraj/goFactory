package main

import "fmt"

func main() {
	fmt.Println(div(3, 0))
	fmt.Println(div(10, 5))
	// demPanic()
}

func div(x, y int) int {
	defer func() {
		fmt.Println(recover())
	}()
	ans := x / y

	return ans
}

// func demPanic() {
// 	defer func() {
// 		fmt.Println(recover())
// 	}()
// 	panic("Panic")
// }
