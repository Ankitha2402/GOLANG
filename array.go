package main

import "fmt"


func sumArray(arr [5]int) int {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return sum
}

func main() {
	
	arr := [5]int{1, 2, 3, 4, 5}
	result := sumArray(arr)
	fmt.Println("The sum of the array elements is:", result)
}
