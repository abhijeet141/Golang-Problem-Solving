package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func countDigits(num int) int {
	cnt := 0
	for num > 0 {
		cnt += 1
		num /= 10
	}
	return cnt
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
	res := countDigits(inputInt)
	log.Println(res)
}
