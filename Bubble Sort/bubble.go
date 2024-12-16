package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		didSwap := 0
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
				didSwap = 1
			}
		}
		if didSwap == 0 {
			break
		}
	}
	return arr
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter numbers seperated by space: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	parts := strings.Split(input, " ")
	var nums []int
	for _, value := range parts {
		num, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}
		nums = append(nums, num)
	}
	arr := bubbleSort(nums)
	for _, value := range arr {
		fmt.Println(value)
	}
}