package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func palindrome(num int) bool {
	res := 0
	x := num
	if num < 0 {
		return false
	}
	for num > 0 {
		rem := num % 10
		res = res*10 + rem
		num /= 10
	}
	if res == x {
		return true
	} else {
		return false
	}
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Number: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	inputInt, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	res := palindrome(inputInt)
	if res {
		log.Println("Palindrome Number")
	} else {
		log.Println("Not Palindrome Number")
	}
}
