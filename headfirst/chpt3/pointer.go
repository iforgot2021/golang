package main

import (
	"fmt"
	"reflect"
)

func main() {
	var myInt int
	fmt.Println(reflect.TypeOf(&myInt))
	var myFloat float64
	fmt.Println(reflect.TypeOf(&myFloat))
	var myBool bool
	fmt.Println(reflect.TypeOf(&myBool))

	myInt = 123

	myIntPoint := &myInt
	fmt.Println(*myIntPoint)

	myFloat = 12.0
	myFloatPoint := &myFloat
	fmt.Println(*myFloatPoint)

	fmt.Println(*createPointer())
}

func createPointer() *float64 {
	myFloat := 98.5
	myFloatPointer := &myFloat
	return myFloatPointer
}
