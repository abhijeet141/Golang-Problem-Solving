package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func selectionSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		mini := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[mini] {
				mini = j
			}
		}
		temp := arr[mini]
		arr[mini] = arr[i]
		arr[i] = temp
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
	arr := selectionSort(nums)
	for _, value := range arr {
		fmt.Println(value)
	}
}
