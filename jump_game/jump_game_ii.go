package main

func Jump(nums []int) int {
	n := len(nums)
	counter := 0
	jumps := 0
	for i := 1; i < n-1; i++ {
		nums[i] = Max(nums[i]+i, nums[i-1])
	}
	for counter < n-1 {
		jumps++
		counter = nums[counter]
	}
	return jumps
}
