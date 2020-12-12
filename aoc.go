package main

import (
	"expensereport"
	"fmt"
	"log"
	"navigation"
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
	// Day 3
	slope := readinput.GetStrings(path + "/input/day3/slope.in")
	fmt.Println("3.1: ", navigation.Traverse(slope, 3, 1, false))
	trajectories := [][]int{{1, 1, 0}, {3, 1, 0}, {5, 1, 0}, {7, 1, 0}, {1, 2, 0}}
	for _, val := range trajectories {
		val[2] = navigation.Traverse(slope, val[0], val[1], false)
	}
	product := 1
	for _, val := range trajectories {
		product = product * val[2]
	}
	fmt.Println("3.2: ", product)
}
