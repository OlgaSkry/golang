package main

import "fmt"

func main() {
	numCh := make(chan []int)
	numSumCh := make(chan int)

	go SumWorker(numCh, numSumCh)
	numCh <- []int{10, 10, 10}

	res := <-numSumCh // 30

	fmt.Println(res)

}

func SumWorker(numCh chan []int, numSumCh chan int) {
	for nums := range numCh {
		numSumCh <- sum(nums)
	}
}

func sum(nums []int) int {
	s := 0
	for _, n := range nums {
		s += n
	}
	return s
}
