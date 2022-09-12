package main

import "fmt"

func main() {
	var name string = "chandresh Raju"
	const pi float64 = 3.141592653589793
	win := true
	num := 10
	fmt.Println(len(name))

	fmt.Println(name + " is chill")
	fmt.Printf("%.3f \n", pi)
	fmt.Printf("%T \n", pi)
	fmt.Printf("%t \n", win)
	fmt.Printf("%d \n", num)

	fmt.Printf("%b \n", 3)
	fmt.Printf("%c \n", 97)

	fmt.Printf("%x \n", 15)
	fmt.Printf("%e \n", pi)
}
