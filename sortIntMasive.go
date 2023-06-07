package main

import (
	"fmt"
)

func main() {
	nums := []int{0, 2, 3, 4, 6, 8, 9}

	s := getSortedMassive(nums)
	fmt.Println(s)
}

func getSortedMassive(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	if len(nums) == 1 {
		return []string{fmt.Sprintf("%d", nums[0])}
	}

	res := make([]string, 0)
	start := nums[0]
	end := nums[0]

	for i := 1; i < len(nums); i++ {
		end = nums[i-1]
		if nums[i]-end > 1 {
			add(start, end, &res)
			start = nums[i]
		}
	}

	end = nums[len(nums)-1]
	add(start, end, &res)

	return res
}

func add(start, end int, result *[]string) {
	if start == end {
		*result = append(*result, fmt.Sprintf("%d", start))
	} else {
		*result = append(*result, fmt.Sprintf("%d->%d", start, end))
	}
}
