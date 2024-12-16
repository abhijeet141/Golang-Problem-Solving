package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func reverseNum(num int) {
	res := 0
	sign := 1
	if num < 0 {
		sign = -1
		num = -num
	}
	for num > 0 {
		rem := num % 10
		res = res*10 + rem
		num /= 10
	}
	res = res * sign
	if res < math.MinInt32 || res > math.MaxInt32 {
		log.Print(0)
	}
	log.Println(res)
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
	reverseNum(inputInt)
}
