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
	lines := strings.Split(string(fileBytes), "\n")
	var lineValues []string
	for _, i := range lines {
		if len(i) > 0 {
			lineValues = append(lineValues, i)
		}
	}
	return lineValues
}

// GetInts reads ints to a slice
func GetInts(fileName string) []int {
	stringValues := GetStrings(fileName)
	var intValues = []int{}
	for _, i := range stringValues {
		j, _ := strconv.Atoi(i)
		intValues = append(intValues, j)
	}
	return intValues
}
