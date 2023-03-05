package main

func CanReach(arr []int, start int) bool {
	n := len(arr)
	queue := []int{start}
	seen := make(map[int]bool)
	for len(queue) > 0 {
		i := queue[0]
		queue = queue[1:]
		if arr[i] == 0 {
			return true
		}
		_, ok := seen[i]
		if ok {
			continue
		}
		seen[i] = true
		if i-arr[i] >= 0 {
			queue = append(queue, i-arr[i])
		}
		if i+arr[i] < n {
			queue = append(queue, i+arr[i])
		}
	}
	return false
}
