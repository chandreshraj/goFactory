package main

import "fmt"

func main() {
	// var age int = 18

	// if age > 18 {
	// 	fmt.Println("you can vote")
	// } else {
	// 	fmt.Println("No, you can't vote")
	// }

	age := 18

	switch age {
	case 18:
		fmt.Println("something.... 18")
	default:
		fmt.Println("Default")
	}

}
