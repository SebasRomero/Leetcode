package main

import "fmt"

func getConcatenation(nums []int) []int {
	nums = append(nums, nums...)
	return nums
}

func main() {
	var nums = []int{1, 2, 1}
	fmt.Println(getConcatenation(nums))
}
