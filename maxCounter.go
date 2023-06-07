package main

import "fmt"

func main() {
	test1 := []int{1, 2, 3, 2, 1, 2, 2}
	test2 := []int{1, 2, 3, 3, 1, 2}

	result, result1 := getMaxCount(test1)
	result2, result3 := getMaxCount(test2)

	fmt.Println(result, "", result1)
	fmt.Println(result2, "", result3)
}

func getMaxCount(massive []int) (int, int) {
	numberCounts := make(map[int]int)
	max := 0
	index := 0
	for _, num := range massive {
		if _, ok := numberCounts[num]; ok {
			numberCounts[num] += 1
		} else {
			numberCounts[num] = 1
		}
	}

	for key, value := range numberCounts {
		if value > max {
			max = value
			index = key
		}
	}
	return index, max
}
