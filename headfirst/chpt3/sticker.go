package main

import "fmt"

func main() {
	var myInt int64
	var myIntPointer *int64
	myInt = 42
	myIntPointer = &myInt
	fmt.Println(*myIntPointer)
}
