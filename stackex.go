package main

import "fmt"

func main() {
	msg1 := "()"

	msg2 := "()[]{}"

	msg3 := "(]"

	check1 := isValid(msg1)
	fmt.Println(check1)
	check2 := isValid(msg2)
	fmt.Println(check2)
	check3 := isValid(msg3)
	fmt.Println(check3)

}

func isValid(s string) bool {
	pairs := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	stack := []rune{}

	for _, x := range s {
		if _, ok := pairs[x]; ok {
			stack = append(stack, x)
		} else if len(stack) == 0 || pairs[stack[len(stack)-1]] != x {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
