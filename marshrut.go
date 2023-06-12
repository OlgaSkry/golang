//Есть массив данных, содержащий маршруты между двумя населенными пунктами вида:
//struct path {
//    char* From;
//    char* To;
//} для C++
//Необходимо написать программу, которая формирует маршрут, включающий посещение всех
//городов из списка.
//Гарантировано, что маршруты не содержат разрывов и закольцованностей.
//Например,
//Исходные данные: ((‘Москва’, ‘Тюмень’), (‘Тюмень’, ‘Сочи’), (‘Ростов-на-Дону’, ‘Москва’))
//Выходные данные: (‘Ростов-на-Дону’, ‘Москва’, ‘Тюмень’, ‘Сочи’)

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
	res, err := createRoute(len(listOfPaths), listOfPaths)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}

func createRoute(lenOfList int, massive []Path) (string, error) {
	//checkLenght
	if lenOfList == 0 {
		return "", errors.New("Invalid data, less than 1 char")
	}

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
	var err error

	for _, k := range massive {
		err = checkValidity(k)
		if err == nil {
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

func checkValidity(path Path) error {
	err1 := checkLetter(path)
	err2 := checkSameTown(path)

	if err1 != nil {
		return err1
	} else {
		return err2
	}
}

func checkSameTown(path Path) error {
	if path.From == path.To {
		return errors.New("Invalid data, same town in the path: " + fmt.Sprintf("%v", &path))
	}
	return nil
}

func checkLetter(path Path) error {
	var IsLetter = regexp.MustCompile(`^[a-zA-Z0-9.,]+$`).MatchString

	if !IsLetter(path.From) || !IsLetter(path.To) {
		return errors.New("There is a wrong simbol: " + fmt.Sprintf("%v", &path))
	}
	return nil
}

//O(3*n)
//for future refactor
//todo validation in a separate module
//todo instead of contains method we can use map[string]struct{} like a set
//todo separate creation of map and serialization to string by logic blocks
//todo check with business where data comes from - library/custom controller/api
