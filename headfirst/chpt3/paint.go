package main

import (
	"fmt"
	"log"
)

func main() {
	var width, height, area float64
	width, height = 4.2, 3.0
	area, err := paintNeeded(width, height)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.2f liters needed\n", area)

	// another
	width, height = 5.2, 3.5
	area, err = paintNeeded(width, height)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.2f liters needed\n", area)
}

func paintNeeded(width, height float64) (float64, error) {
	if width <= 0 {
		return 0, fmt.Errorf("a width of %.2f is invalid", width)
	}
	if height <= 0 {
		return 0, fmt.Errorf("a height of %.2f is invalid", height)
	}
	return width * height / 10.0, nil
}
