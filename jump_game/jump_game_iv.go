package main

import "container/list"

func MinJumps(arr []int) int {
	n := len(arr)
	positions := make(map[int][]int)
	for i, num := range arr {
		positions[num] = append(positions[num], i)
	}
	q := list.New()
	q.PushBack([]int{0, 0})
	seen := make([]bool, n)
	for q.Len() > 0 {
		head := q.Front()
		i := head.Value.([]int)[0]
		jumps := head.Value.([]int)[1]
		q.Remove(head)
		if i == n-1 {
			return jumps
		}
		seen[i] = true
		for _, pos := range positions[arr[i]] {
			if !seen[pos] {
				q.PushBack([]int{pos, jumps + 1})
			}
		}
		delete(positions, arr[i])
		if i > 0 && !seen[i-1] {
			q.PushBack([]int{i - 1, jumps + 1})
		}
		if !seen[i+1] {
			q.PushBack([]int{i + 1, jumps + 1})
		}
	}
	return -1
}
