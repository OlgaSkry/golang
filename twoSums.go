package main

import "fmt"

func main() {
	test1 := []int{2, 7, 11, 15}
	target1 := 9

	test2 := []int{2, 3, 4}
	target2 := 7

	test3 := []int{3, 3}
	target3 := 6

	check1 := getIndexes(test1, target1)
	fmt.Println(check1)
	check2 := getIndexes(test2, target2)
	fmt.Println(check2)
	check3 := getIndexes(test3, target3)
	fmt.Println(check3)
}

func getIndexes(massive []int, target int) []int {
	indexMap := make(map[int]int)

	for i, num := range massive {
		if recIndex, ok := indexMap[target-num]; ok {
			return []int{recIndex, i}
		} else {
			indexMap[num] = i
		}
	}
	return []int{}

}
