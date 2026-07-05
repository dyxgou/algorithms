package main

import (
	"log/slog"
)

func abs(n int) int {
	if n < 0 {
		n = -n
	}

	return n
}

func firstMissingPositive(nums []int) int {
	for i, n := range nums {
		if n >= 0 {
			continue
		}

		nums[i] = 0
	}

	for i, n := range nums {
		if n == 0 {
			nums[i] = len(nums) + 1
		}

		idx := abs(n) - 1
		if idx < 0 || idx > len(nums)-1 || nums[idx] < 0 {
			continue
		}

		nums[idx] = -nums[idx]
		slog.Info("third", "nums", nums, "n", n)
	}

	for n := 1; n <= len(nums); n++ {
		if nums[n-1] > 0 {
			return n
		}
	}

	return len(nums) + 1
}

func main() {
	n := firstMissingPositive([]int{
		98, 93, 95, 10, 91, 4, 90, 88, 56, 84, 65, 62, 83, 80, 78, 60, 73, 77, 76, 29, 63, 12, 57, 17, 69, 68, 50, 11, 31, 33, 8, 42, 38, 7, 0, 37, 48, 26, 20, 44, 46, 43, 52, 51, 47, 18, 49, 58, 2, 39, 30, 81, 22, 55, 36, 40, 15, 27, 21, 32, 64, 41, 53, 19, 28, 24, 9, 25, 3, 59, 66, 82, 61, 70, 23, 34, 71, 54, 74, -1, 1, 45, 14, 79, 5, 35, 13, 72, 75, 85, 87, 6, 16, 86, 67, 89, 94, 92, 96, 97,
	})

	slog.Info("first missing positive", "n", n)
}
