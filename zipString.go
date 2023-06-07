package main

import "fmt"

func main() {
	stringBefore := []string{"a", "a", "b", "b"}
	res := zipStr(stringBefore)
	fmt.Println(res)
}

func zipStr(strToProccess []string) string {
	var char string
	count := 1
	var result string
	for i := 0; i < len(strToProccess)-1; i++ {
		char = strToProccess[i]
		if strToProccess[i] == strToProccess[i+1] {
			count += 1
		} else {
			result += fmt.Sprint(strToProccess[i], count)
			count = 1
		}
	}
	result += fmt.Sprint(char, count)
	return result
}
