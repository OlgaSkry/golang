package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type Path struct {
	From string
	To   string
}

func main() {
	var listOfPaths = []Path{
		{
			From: "msc",
			To:   "vor",
		},
		{
			From: "vor",
			To:   "cheb",
		},
		{
			From: "ros",
			To:   "msc",
		},
	}
	res, err := createMarhsrut(len(listOfPaths), listOfPaths)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

}

func createMarhsrut(lenOfList int, massive []Path) (string, error) {
	result := make([]string, lenOfList)

	mapOfPaths, listOfTos, err := processDataForValidation(massive)

	if err != nil {
		return "", err
	}

	//get the head of linked list

	var head string

	for _, k := range massive {
		s := contains(listOfTos, k.From)
		if !s {
			head = k.From
		}
	}

	result = append(result, head)

	for i := 0; i <= lenOfList; i++ {
		result = append(result, mapOfPaths[head])
		head = mapOfPaths[head]
	}

	result1 := strings.Join(result, " ")

	return result1, nil
}

//go doesn't have contains for slice structure

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

//validation

func processDataForValidation(massive []Path) (map[string]string, []string, error) {

	mapOfPaths := make(map[string]string)
	var listOfTos []string
	var isValid bool
	var err error

	for _, k := range massive {
		isValid, err = checkValidity(k)
		if isValid {
			if _, ok := mapOfPaths[k.From]; !ok {
				mapOfPaths[k.From] = k.To
			} else {
				return mapOfPaths, listOfTos, errors.New("Invalid data, same town in list of paths From: " + fmt.Sprintf("%v", k.From))
			}
			isLoop := contains(listOfTos, k.To)
			if !isLoop {
				listOfTos = append(listOfTos, k.To)
			} else {
				return mapOfPaths, listOfTos, errors.New("Invalid data, same town in list of paths To: " + fmt.Sprintf("%v", k.To))
			}

		}
	}
	return mapOfPaths, listOfTos, err
}

func checkValidity(path Path) (bool, error) {
	isValidByLen, err1 := checkMoreThanOne(path)
	isValidBySymbol, err2 := checkLetter(path)
	isValidWithoutDublicate, err3 := checkSameTown(path)

	isValid := isValidByLen && isValidBySymbol && isValidWithoutDublicate

	if err1 != nil {
		return isValid, err1
	} else if err2 != nil {
		return isValid, err2
	} else {
		return isValid, err3
	}
}

func checkSameTown(path Path) (bool, error) {
	if path.From == path.To {
		return false, errors.New("Invalid data, same town in the path: " + fmt.Sprintf("%v", &path))
	}
	return true, nil
}

func checkMoreThanOne(path Path) (bool, error) {
	if len(path.From) < 1 || len(path.To) < 1 {
		return false, errors.New("Invalid data, less than 1 char: " + fmt.Sprintf("%v", &path))
	}
	return true, nil
}

func checkLetter(path Path) (bool, error) {
	var IsLetter = regexp.MustCompile(`^[a-zA-Z0-9.,]+$`).MatchString

	if !IsLetter(path.From) || !IsLetter(path.To) {
		return false, errors.New("There is a wrong simbol: " + fmt.Sprintf("%v", &path))
	}
	return true, nil
}
