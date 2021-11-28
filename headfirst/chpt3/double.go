package main

import "fmt"

func Double(num *int64) {
	*num *= 2
}

func main() {
	var amount int64 = 6
	Double(&amount)
	fmt.Println(amount)
}
