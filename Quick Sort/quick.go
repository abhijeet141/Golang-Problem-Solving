package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Partition(arr []int, low, high int) int {
	pivot := arr[low]
	i := low
	j := high
	for i < j {
		for arr[i] <= pivot && i <= high-1 {
			i++
		}
		for arr[j] > pivot && j >= low+1 {
			j--
		}
		if i < j {
			temp := arr[i]
			arr[i] = arr[j]
			arr[j] = temp
		}
	}
	temp := arr[low]
	arr[low] = arr[j]
	arr[j] = temp
	return j
}
func quickSort(arr []int, low, high int) {
	if low >= high {
		return
	}
	partitionIndex := Partition(arr, low, high)
	quickSort(arr, low, partitionIndex-1)
	quickSort(arr, partitionIndex+1, high)
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
	low := 0
	high := len(nums) - 1
	quickSort(nums, low, high)
	fmt.Print(nums)
}
