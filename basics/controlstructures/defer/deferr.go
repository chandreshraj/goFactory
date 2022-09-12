package main

import "fmt"

func main() {
	defer FirstRun()
	SecondRun()
}

func FirstRun() {
	fmt.Println("Exe 1st")
}

func SecondRun() {
	fmt.Println("Exe 2nd")
}
