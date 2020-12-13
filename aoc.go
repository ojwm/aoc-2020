package main

import (
	"document"
	"expensereport"
	"fmt"
	"log"
	"navigation"
	"os"
	"passwordutil"
	"readinput"
)

var path string
var err error

func main() {
	path, err = os.Getwd()
	if err != nil {
		log.Println(err)
	}
	day1()
	day2()
	day3()
	day4()
}

func day1() {
	expenses := readinput.GetInts(path+"/input/day1/expenses.in", true)
	sum := 2020
	fmt.Println("1.1: ", expensereport.Find(expenses, sum, 2))
	fmt.Println("1.2: ", expensereport.Find(expenses, sum, 3))
}

func day2() {
	passwords := readinput.GetStrings(path+"/input/day2/passwords.in", true)
	fmt.Println("2.1: ", passwordutil.Validate(passwords, "sled"))
	fmt.Println("2.2: ", passwordutil.Validate(passwords, "toboggan"))
}

func day3() {
	slope := readinput.GetStrings(path+"/input/day3/slope.in", true)
	fmt.Println("3.1: ", navigation.Traverse(slope, 3, 1, false))
	trajectories := [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	product := 1
	for _, val := range trajectories {
		product *= navigation.Traverse(slope, val[0], val[1], false)
	}
	fmt.Println("3.2: ", product)
}

func day4() {
	documents := readinput.GetStrings(path+"/input/day4/documents.in", false)
	cleanedDocuments := document.CleanDocuments(documents)
	validFormatDocuments := document.ValidateFormat(cleanedDocuments)
	fmt.Println("4.1: ", len(validFormatDocuments))
	validContentDocuments := document.ValidateContent(validFormatDocuments)
	fmt.Println("4.2: ", len(validContentDocuments))
}
