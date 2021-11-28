package main

import "fmt"

func main() {
	fmt.Println(status(59.9))
	fmt.Println(status(73.2))
}

func status(grade float64) string {
	if grade >= 60.0 {
		return "passing"
	}
	return "failing"
}
