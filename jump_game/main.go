package main

import "fmt"

func MaxSlice(slice []int) int {
	currMax := slice[0]
	for _, num := range slice {
		currMax = Max(currMax, num)
	}
	return currMax
}

func Min(a int, b int) int {
	if a <= b {
		return a
	}
	return b
}

func Max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

func TestBool(myFunc func(nums []int) bool, numsList [][]int, answers []bool) {
	for i, nums := range numsList {
		result := myFunc(nums)
		if result == answers[i] {
			fmt.Printf("Test case %d - Pass\n", i)
		} else {
			fmt.Printf("Test case %d - Fail\n", i)
		}
	}
}

func TestBool2(myFunc func(nums []int, start int) bool, numsList [][]int, starts []int, answers []bool) {
	for i, nums := range numsList {
		result := myFunc(nums, starts[i])
		if result == answers[i] {
			fmt.Printf("Test case %d - Pass\n", i)
		} else {
			fmt.Printf("Test case %d - Fail\n", i)
		}
	}
}

func TestInt(myFunc func(nums []int) int, numsList [][]int, answers []int) {
	for i, nums := range numsList {
		result := myFunc(nums)
		if result == answers[i] {
			fmt.Printf("Test case %d - Pass\n", i)
		} else {
			fmt.Printf("Test case %d - Fail\n", i)
		}
	}
}

func TestInt2(myFunc func(nums []int, d int) int, numsList [][]int, ds []int, answers []int) {
	for i, nums := range numsList {
		result := myFunc(nums, ds[i])
		if result == answers[i] {
			fmt.Printf("Test case %d - Pass\n", i)
		} else {
			fmt.Printf("Test case %d - Fail\n", i)
		}
	}
}

func main() {
	fmt.Println("Jump Game 1:")
	numsList1 := [][]int{{2, 3, 1, 1, 4}}
	numsList1 = append(numsList1, []int{3, 2, 1, 0, 4})
	numsList1 = append(numsList1, []int{0})
	answers1 := []bool{true, false, true}
	TestBool(CanJump, numsList1, answers1)

	fmt.Println("Jump Game 2:")
	numsList2 := [][]int{{2, 3, 1, 1, 4}}
	numsList2 = append(numsList2, []int{2, 3, 0, 1, 4})
	numsList2 = append(numsList2, []int{0})
	numsList2 = append(numsList2, []int{1, 2, 3, 4, 42, 5, 6})
	answers2 := []int{2, 2, 0, 3}
	TestInt(Jump, numsList2, answers2)

	fmt.Println("Jump Game 3:")
	numsList3 := [][]int{{4, 2, 3, 0, 3, 1, 2}}
	numsList3 = append(numsList3, []int{4, 2, 3, 0, 3, 1, 2})
	numsList3 = append(numsList3, []int{3, 0, 2, 1, 2})
	starts := []int{5, 0, 2}
	answers3 := []bool{true, true, false}
	TestBool2(CanReach, numsList3, starts, answers3)

	fmt.Println("Jump Game 4:")
	numsList4 := [][]int{{100, -23, -23, 404, 100, 23, 23, 23, 3, 404}}
	numsList4 = append(numsList4, []int{7})
	numsList4 = append(numsList4, []int{7, 6, 9, 6, 9, 6, 9, 7})
	answers4 := []int{3, 0, 1}
	TestInt(MinJumps, numsList4, answers4)

	fmt.Println("Jump Game 5:")
	numsList5 := [][]int{{6, 4, 14, 6, 8, 13, 9, 7, 10, 6, 12}}
	numsList5 = append(numsList5, []int{3, 3, 3, 3, 3})
	numsList5 = append(numsList5, []int{7, 6, 5, 4, 3, 2, 1})
	ds5 := []int{2, 3, 1}
	answers5 := []int{4, 1, 7}
	TestInt2(MaxJumps, numsList5, ds5, answers5)
}
