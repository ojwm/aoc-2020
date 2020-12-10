package readinput

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// GetStrings reads strings to a slice
func GetStrings(fileName string) []string {
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}
	return strings.Split(string(fileBytes), "\n")
}

// GetInts reads ints to a slice
func GetInts(fileName string) []int {
	stringValues := GetStrings(fileName)
	var intValues = []int{}
	for _, i := range stringValues {
		if len(i) > 0 {
			j, _ := strconv.Atoi(i)
			intValues = append(intValues, j)
		}
	}
	return intValues
}
