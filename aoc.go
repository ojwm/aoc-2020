package main

import (
	"expensereport"
	"fmt"
	"log"
	"os"
	"passwordutil"
	"readinput"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	// Day 1
	expenses := readinput.GetInts(path + "/input/day1/expenses.in")
	sum := 2020
	fmt.Println("1.1: ", expensereport.Find(expenses, sum, 2))
	fmt.Println("1.2: ", expensereport.Find(expenses, sum, 3))
	// Day 2
	passwords := readinput.GetStrings(path + "/input/day2/passwords.in")
	fmt.Println("2.1: ", passwordutil.Validate(passwords, "sled"))
	fmt.Println("2.2: ", passwordutil.Validate(passwords, "toboggan"))
}
