package main

func CanJump(nums []int) bool {
	n := len(nums)
	currMax := 0
	for i, num := range nums[:len(nums)-1] {
		if i > currMax {
			return false
		}
		currMax = Max(currMax, i+num)
		if currMax >= n-1 {
			return true
		}
	}
	return currMax >= n-1
}
