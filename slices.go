package main

import "fmt"
func maxInSlice(nums []int) (int, error) {
	if len(nums) == 0 {
		return 0, fmt.Errorf("slice is empty")
	}
	max := nums[0]
	for _, value := range nums {
		if value > max {
			max = value
		}
	}
	return max, nil
}
func main() {
	nums := []int{8,9,5,2,100}
	maxValue, err := maxInSlice(nums)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("The maximum value in the slice is:", maxValue)
	}
}
