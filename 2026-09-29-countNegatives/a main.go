package main

import "fmt"

func countNegatives(nums []int) int {
	count := 0
	for _, v := range nums {
		if v < 0 {
			count++
		}
	}
	return count
}

func main() {
	firstSlice := []int{3, -1, 0, -7, 5}
	secondSlice := []int{1, 2, 3}
	thirdSlice := []int{}
	fmt.Println(countNegatives(firstSlice))
	fmt.Println(countNegatives(secondSlice))
	fmt.Println(countNegatives(thirdSlice))
}
