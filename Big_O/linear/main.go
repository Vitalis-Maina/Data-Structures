package main

import "fmt"

func main() {

	nums := []int{23, 45, 56, 2, 34, 12, 15, 64}

	for i := 0; i < len(nums); i++ {
		square := 0
		square = nums[i] * nums[i]
		fmt.Printf("square of %d : %d\n", nums[i], square)
	}

}
