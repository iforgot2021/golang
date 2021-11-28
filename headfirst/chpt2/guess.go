package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var success = false
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100) + 1
	fmt.Println("please input a number")
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < 10; i++ {
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		num, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if num == target {
			success = true
			break
		} else if num > target {
			fmt.Println("too higher")
		} else {
			fmt.Println("too lower")
		}
	}
	if !success {
		fmt.Println("sorry, you failed")
	}
}
