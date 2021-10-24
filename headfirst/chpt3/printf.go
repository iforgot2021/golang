package main

import "fmt"

func main() {
	fmt.Printf("%d\n", 123)
	fmt.Printf("%f\n", 123.0)
	fmt.Printf("%s\n", "123")
	fmt.Printf("%c\n", 65)
	fmt.Printf("%t\n", true)
	fmt.Printf("%T\n", [3]int{0, 1, 2})
	fmt.Printf("%v\n", [2]int{0, 1})
	fmt.Printf("%#v\n", 123)
	fmt.Printf("%%\n")
}
