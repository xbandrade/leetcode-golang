package main

import "sort"

// #1340 https://leetcode.com/problems/jump-game-v/

func MaxJumps(arr []int, d int) int {
	n := len(arr)
	dp := make([]int, n)
	sortedIndices := make([]int, n)
	for i := range sortedIndices {
		sortedIndices[i] = i
	}
	sort.Slice(sortedIndices, func(i, j int) bool {
		return arr[sortedIndices[i]] < arr[sortedIndices[j]]
	})
	for _, i := range sortedIndices {
		currMax := 1
		left := Max(0, i-d) - 1
		right := Min(i+d+1, n)
		for k := i - 1; k > left; k-- {
			if arr[k] >= arr[i] {
				break
			}
			currMax = Max(dp[k]+1, currMax)
		}
		for k := i + 1; k < right; k++ {
			if arr[k] >= arr[i] {
				break
			}
			currMax = Max(dp[k]+1, currMax)
		}
		dp[i] = currMax
	}
	return MaxSlice(dp)
}
