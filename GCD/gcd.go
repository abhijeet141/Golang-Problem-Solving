package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func gcd(a, b int) {
	for a > 0 && b > 0 {
		if a > b {
			a = a % b
		} else {
			b = b % a
		}
	}
	if a == 0 {
		log.Println(b)
	}
	log.Println(a)
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Number1: ")
	input1, _ := reader.ReadString('\n')
	input1 = strings.TrimSpace(input1)
	inputInt1, err1 := strconv.Atoi(input1)
	fmt.Print("Enter Number2: ")
	input2, _ := reader.ReadString('\n')
	input2 = strings.TrimSpace(input2)
	inputInt2, err2 := strconv.Atoi(input2)
	if err1 != nil || err2 != nil {
		panic(err1)
	}
	gcd(inputInt1, inputInt2)
}
