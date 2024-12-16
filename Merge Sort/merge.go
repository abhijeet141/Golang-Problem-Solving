package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func merge(nums []int, low, mid, high int) {
	var temp []int
	left := low
	right := mid + 1
	for left <= mid && right <= high {
		if nums[left] <= nums[right] {
			temp = append(temp, nums[left])
			left++
		} else {
			temp = append(temp, nums[right])
			right++
		}
	}
	for left <= mid {
		temp = append(temp, nums[left])
		left++
	}
	for right <= high {
		temp = append(temp, nums[right])
		right++
	}
	for i := low; i <= high; i++ {
		nums[i] = temp[i-low]
	}
}
func mergeSort(nums []int, low, high int) {
	if low >= high {
		return
	}
	mid := (low + high) / 2
	mergeSort(nums, low, mid)
	mergeSort(nums, mid+1, high)
	merge(nums, low, mid, high)
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
	mergeSort(nums, low, high)
	// for _, value := range nums {
	// 	fmt.Println(value)
	// }
	fmt.Println(nums)
}
