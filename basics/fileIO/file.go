package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	file, error := os.Create("Sample.txt")
	if error != nil {
		log.Fatal(error)
	}
	file.WriteString("Something")
	defer file.Close()

	stream, error := ioutil.ReadFile("./Sample.txt")
	if error != nil {
		log.Fatal(error)
	}

	s1 := string(stream)
	fmt.Println(s1)
}
