package readinput

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// getLines reads lines to a slice
func getLines(fileName string) []string {
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}
	return strings.Split(string(fileBytes), "\n")
}

// removeBlankValues from a slice
func removeBlankValues(values []string) []string {
	var returnValues []string
	for _, i := range values {
		if len(i) > 0 {
			returnValues = append(returnValues, i)
		}
	}
	return returnValues
}

// GetStrings reads strings to a slice
func GetStrings(fileName string, removeBlanks bool) []string {
	stringValues := getLines(fileName)
	if removeBlanks {
		return removeBlankValues(stringValues)
	}
	return stringValues
}

// GetInts reads ints to a slice
func GetInts(fileName string, removeBlanks bool) []int {
	stringValues := GetStrings(fileName, removeBlanks)
	var intValues = []int{}
	for _, i := range stringValues {
		j, _ := strconv.Atoi(i)
		intValues = append(intValues, j)
	}
	return intValues
}
