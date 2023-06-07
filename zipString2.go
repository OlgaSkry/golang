package main

import (
	"errors"
	"fmt"
	"regexp"
)

func main() {
	testString1 := []string{"a", "a", "b", "b"}
	testString2 := []string{"a"}
	testString3 := []string{"a", "1", "b", ","}

	var result1 string
	var result2 string
	var result3 string

	result1, err := getZipString(testString1)
	result2, err1 := getZipString(testString2)
	result3, err2 := getZipString(testString3)

	printResult(result1, err)
	printResult(result2, err1)
	printResult(result3, err2)
}

func getZipString(s []string) (string, error) {

	var strToProcess string
	count := 1
	k := s[0]

	var IsLetter = regexp.MustCompile(`^[a-zA-Z0-9.,]+$`).MatchString

	for i := 0; i < len(s)-1; i++ {
		k = s[i]
		if !IsLetter(k) {
			return strToProcess, errors.New("There is a wrong simbol: " + k)
		}
		if k == string(s[i+1]) {
			count += 1
		} else {
			strToProcess += fmt.Sprint(k, count)
			count = 1
		}
	}
	strToProcess += fmt.Sprint(k, count)
	return strToProcess, nil
}

func printResult(s string, err error) {
	if err == nil {
		fmt.Println(s)
	} else {
		fmt.Print(err)
	}
}
