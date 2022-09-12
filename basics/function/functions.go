package main

import "fmt"

func main() {
	x, y := 5, 6
	fmt.Println(add(x, y))

	fmt.Println(addAll(10, 20, 30, 40, 50))

}

func addAll(args ...int) int {
	sum := 0
	for _, value := range args {
		sum += value
	}
	return sum
}

func add(x, y int) int {
	return x + y
}
