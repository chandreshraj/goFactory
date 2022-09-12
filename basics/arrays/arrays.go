package main

import "fmt"

func main() {
	// var EvenNum [5]int
	// EvenNum[0] = 0
	// EvenNum[1] = 1
	// EvenNum[2] = 2
	// EvenNum[3] = 3
	// EvenNum[4] = 4

	// // fmt.Println(EvenNum[3])

	// for i, value := range EvenNum {
	// 	fmt.Println(EvenNum[i], value)
	// }

	//SLICING
	// numSlice := []int{0, 9, 8, 7, 6, 5}
	// sliced := numSlice[0:]
	// fmt.Println(sliced)

	//Create Empty Slice & copy slices
	numSlice := []int{0, 9, 8, 7, 6, 5}

	slice2 := make([]int, 6, 10)
	copy(slice2, numSlice)

	fmt.Println(slice2)

	// Append slices

	slice3 := append(numSlice, 3, 0, -1)
	fmt.Println(slice3)
}
