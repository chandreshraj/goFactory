package main

import "fmt"

func main() {
	//FOR
	fmt.Println("For Loop")
	for i := 1; i <= 3; i++ {
		fmt.Println(i)
	}
	//while
	i := 1
	fmt.Println("While Using FOR")
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	//FOR inside FOR
	for i := 1; i < 5; i++ {
		for j := 1; j < i; j++ {
			fmt.Println("*")
		}
		fmt.Println(" ")
	}
}
